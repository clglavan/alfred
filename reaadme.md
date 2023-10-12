Based on your Azure access run this binary to gather info about TBD

It only executes
- `az account list --output table` for showing and `az account list --output json` to get name and ID of subscription
- `az aks list --subscription sub.ID --output json` inside a foreach looping through subscriptions and passing it
- then just going through the json

it's not fast, it can be done better.
