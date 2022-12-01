# Prompts

1. Design and provide the input and output schema for your API. Some things to consider:
    - How will the schema take into consideration if the person being provided is not on English
wikipedia.
    - How will you account for missing “short description” in content.

    **A:** Considerations behind decisions regarding the design of the API and schema can be found in [considerations.md](./considerations.md).
    Information on the schema can be found in the [README.md](../README.md).

2. Please provide your functioning code in a public github repo. One should be able to run your API server
locally and try out with some inputs.

    **A:** Source code for the API can be found in the GitHub repo, [wkyoshida/wmf-task](https://github.com/wkyoshida/wmf-task).

3. Consider this hypothetical scenario: Your API is going to be deployed and made available to the public for
use. What things could you do to keep this API service highly available and reliable? (Think of as many
issues as come to your mind and propose your potential solutions. No code is required for this)

    **A:** Considerations regarding potential General Availability of the API can be found in [considerations.md](./considerations.md).
