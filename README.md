# GitHub Actions Secret Detection Workflow

This repository demonstrates a GitHub Actions workflow that scans pull requests for leaked secrets and credentials.

## How It Works

The workflow consists of two jobs:

1. **Secret Scanning (gitleaks-scan)**:
   - Uses [Gitleaks](https://github.com/gitleaks/gitleaks) to detect potential secrets in the code
   - **Important**: Only scans the current state of the PR branch, not the commit history
   - If secrets are found:
     - Posts a comment on the PR with details about the leaked secrets
     - Uploads a detailed report as an artifact
     - Fails the workflow to prevent merging
   - If no secrets are found, the workflow proceeds to the second job

2. **Additional Checks (second-job)**:
   - Only runs if no secrets are detected in the first job
   - Performs additional security checks (placeholder for your custom checks)
   - Completes the workflow successfully

## Testing the Workflow

The repository includes a `test-secrets.txt` file with fake credentials to demonstrate how the secret detection works. In a real-world scenario, these would trigger the detection mechanism and fail the workflow.

## Usage

1. Create a pull request with changes
2. The workflow automatically runs and scans for secrets
3. If secrets are found, the PR will be flagged and cannot be merged until the secrets are removed
4. If no secrets are found, additional checks will run

## Configuration Details

- The workflow is configured to only scan the current state of files in the PR branch
- It uses `--no-git=true` to prevent scanning git history
- It uses `fetch-depth: 1` to only fetch the latest commit
- A custom `.gitleaks.toml` configuration file is used to:
  - Ignore documentation files like README.md where examples might trigger false positives
  - Allowlist specific patterns that are commonly used in examples but might look like secrets
  - Prevent false positives while still catching real secrets

## Important Note

The fake credentials in `test-secrets.txt` are for demonstration purposes only and should not be used in any real application.
