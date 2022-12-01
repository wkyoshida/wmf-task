# Considerations

## Technologies

- [go-mwclient](https://github.com/cgt/go-mwclient) is listed in the [API client libraries](https://www.mediawiki.org/wiki/API:Client_code#Go) 
resource page. The library uses the 
[version 2](https://www.mediawiki.org/wiki/API:JSON_version_2) 
of the MediaWiki JSON API, which is the version needed.
It also appeared more up-to-date over the other client library listed.
Considering the parameters given in the MediaWiki API example:

```
https://en.wikipedia.org/w/api.php?
action=query
&prop=revisions
&titles=Yoshua_Bengio
&rvlimit=1
&formatversion=2
&format=json
&rvprop=content
```

Their effect are also accounted for in go-mwclient
[here](https://github.com/cgt/go-mwclient/blob/8baad1652addcd819c9f5e7de80328ba350799e9/edit.go#L119-L122)
and [here](https://github.com/cgt/go-mwclient/blob/8baad1652addcd819c9f5e7de80328ba350799e9/core.go#L159-L163).
Use of a MediaWiki API client library was considered to help ease 
some of the groundwork to get started faster.

- [net/http](https://pkg.go.dev/net/http) for the HTTP handling
was considered as an option that is light-weight and also already 
available in the standard library, which seemed best for the purposes of 
quick, simple API implementation. Other popular alternatives, such as [gin-gonic/gin](https://github.com/gin-gonic/gin)
could be considered, but perhaps too heavy. 

- Use of config file, i.e. [env.yaml](../env.yaml), was considered to 
allow for easier modification of server configuration, 
not requiring recompilation of the source code for changes.  

- regex was considered to help parse out the `{{Short description}}`
from the content of a MediaWiki page. Found an already implemented regex
used by [wikimedia/wikipedia-ios](https://github.com/wikimedia/wikipedia-ios/blob/18e9521bd0f04ea2503d180231da180488dda3e1/Wikipedia/Code/Controllers/ShortDescriptionController.swift#L27)
that was leveraged. Seems to also align with guidance on how to write a
[Short_description](https://en.wikipedia.org/wiki/Wikipedia:Short_description#How_to_edit).

## MediaWiki Considerations

- Decision to include the fields in the response schema partly due to:  
  - `pageid`: See note.
  - `title`: See note.
  - `shortdescription`: Field is central to the purpose of the API.
  Inclusion seems fit as the field can be used when the content of a 
  page's short description is needed in a more general-use format.
  - `shortdescriptionraw`: Field is the raw short description still in
  MediaWiki template form. Inclusion seems fit as the field can be used
  for more MediaWiki-specific purposes.
  - `timestamp`: See note.

  **Note:** Field is a transfer from the response returned by the 
  MediaWiki API. Inclusion seems fit as the field represents a particular 
  resource in the MediaWiki instance and can be used to relate back to it.  

  See [README.md](../README.md#response-schema) for schema.

- Decision to make `titles` a query parameter partly due to design choice 
of the MediaWiki API to also use query parameters itself. Following 
same pattern could be familiar for users hitting APIs for both MediaWiki and shortdesc. 

- Given that shortdesc-api makes use of the [MediaWiki API](https://www.mediawiki.org/wiki/API:Query#API_documentation),
decision was made to communicate the same limit with the `titles` parameter in
the schema docs, i.e. the maximum number of `titles` values in a single request
allowed is 50 values. This is the same limit with the `titles` parameter
in the MediaWiki API.

- Redirects for a page title are possible as mentioned in 
[Resolving redirects](https://www.mediawiki.org/wiki/API:Query#Resolving_redirects). 
However: 
  - A user could already know or discover the page title through other 
  means first, e.g. via the 
  [Search API](https://www.mediawiki.org/wiki/API:REST_API/Reference#Search),
  and later call the shortdesc-api.
  - Configuring go-mwclient to request redirects is possible, but not configured by default.
  - Could make sense for `titles` to only directly relate to resources 
  (pages) in the MediaWiki instance, i.e. to not do the work of trying to
  suggest or find the page the user could have meant by their request. 
  (Related to next point)   

- Page not availabe on English instance?
  - Could make sense to not return an entry in the response for the page
  if the page does not exist in the instance.
  - The config file allowing multiple MediaWiki instances, e.g. English and Portuguese Wikipedia, could allow for 
  future enhancements:
    - The API could be enhanced to have an additional query parameter allowing 
    a user to select which instance to make the request for. This would a user
    to make requests for different instances. However, the API would need a
    way to communicate which instances are supported byt the API 
    (either documentation or another API endpoint that lists the instances).  
    - An alternative could be to simply host an instance of the API 
    specific per MediaWiki instance. This would allow each MediaWiki instance 
    to have their own shortdesc-api. 

- Page missing `{{Short description}}`?
  - Could make sense to return the rest of the data for the page normally, 
  yet with the `shortdescription` and `shortdescriptionraw` fields empty. 
  User is now made aware that the page exists, but a short description does 
  not yet exist for it.

## General Availability

- Technology and architecture
  - Investment on dedicated infra, i.e. hardware, or cloud can help 
  the API service have the required resources to serve increased demand
  - Caching can be leveraged as well to ease system burden by not performing 
  full data retrievals/processing for every request
  - Distribute the request processing can help ease the burden on single 
  hardware and also help scale the service:
    - Geographically-distributed hardware for multiple global regions; 
    multiple workers
    - Leverage load-balancers to distribute requests/work 
    - Containerazing the API with a tool like 
    [Docker](https://en.wikipedia.org/wiki/Docker_(software)) can also help 
    the service to scale as it can be spun up quicker in different machines
    - Alongside the above point, container orchestration technologies lik
    [kubernetes](https://github.com/kubernetes/kubernetes) can help deliver
    that scale with containers and manage multiple workers.

- Community engagement
  - Communicate best practices on how users should use the API
  - Monitor for bad actors and, if needed, IP-block or something similar
  - If needed, a rate limit could be enforced

- Good practices by the development team can help keep maintenance of the project: 
  - Testing, e.g. unit, integration, load, etc., to check for code quality. 
    - For instance, have to look to leverage interfaces in the code to 
    properly use [mock](https://pkg.go.dev/github.com/stretchr/testify/mock) 
    in the current unit tests.
  - Set up good CI, e.g. GitHub Actions, Jenkins, CircleCI, AWS CodeBuild, etc. Can help the team iterate quickly and with more confidence 
  on features/functionality.
  - Set up tools for observability, logging, metrics, alerts, uptime 
  monitoring. Can help the team better understand the health of the service
  and be aware of issues/outages that arise.
  - Dedicated SRE, Ops, support team, or support rotation. 
  Dedicated engineering resources for ensuring uptime/operations/reliability 
  can help as well. Development resources can also then become more free to work on new features.
