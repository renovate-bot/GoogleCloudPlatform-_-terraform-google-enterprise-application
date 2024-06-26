# 3. Application Factory phase

## Purpose
The application factory creates application project groups, which contain resources responsible for deployment of a single application within the developer platform.

An overview of the application factory pipeline is shown below.
![Enterprise Application application factory diagram](assets/eab-app-factory.svg)

The application factory creates the following resources as defined in the [`app-group-baseline`](./modules/app-group-baseline/) submodule:

* **Application admin project:** A project dedicated for application administration and management.
* **Application environment projects:** A project for the application for each environment (e.g., development, nonproduction, production).
* **Infrastructure repository:** A Git repository containing the Terraform configuration for the application infrastructure.
* **Application infrastucture pipeline:** A Cloud Build pipeline for deploying the application infrastructure specified as Terraform.


## Usage

### Running Terraform locally

1. The next instructions assume that you are in the `terraform-google-enterprise-application/3-appfactory` folder.

   ```bash
   cd terraform-google-enterprise-application/3-appfactory
   ```

1. Rename `terraform.example.tfvars` to `terraform.tfvars`.

   ```bash
   mv terraform.example.tfvars terraform.tfvars
   ```

1. Update the file with values for your environment.

You can now deploy the into your common folder.

1. Run `init` and `plan` and review the output.

   ```bash
   terraform init -chdir=./apps/cymbal-bank
   terraform plan -chdir=./apps/cymbal-bank
   ```

1. Run `apply`.

   ```bash
   terraform apply -chdir=./apps/cymbal-bank
   ```

If you receive any errors or made any changes to the Terraform config or `terraform.tfvars`, re-run `terraform plan -chdir=./apps/cymbal-bank` before you run `terraform apply -chdir=./apps/cymbal-bank`.
