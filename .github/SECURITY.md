# Security Policy

## Supported Versions

We support the latest release of the Solace Software Event Broker Terraform Provider. Older versions do not receive security fixes.

## Reporting a Vulnerability

**Do not report security vulnerabilities via public GitHub issues.**

We take security seriously. If you discover a security vulnerability in the Solace Software Event Broker Terraform Provider, please report it responsibly.

### How to Report

**GitHub private vulnerability reporting (preferred)**: [Report a vulnerability](https://github.com/SolaceProducts/terraform-provider-solacebroker/security/advisories/new)

**Email**: [support@solace.com](mailto:support@solace.com)

Support triages every report and escalates anything describing a concrete, exploitable vulnerability directly to our security team.

### What to Include

- **Description** of the vulnerability
- **Steps to reproduce** the issue
- **Potential impact** (what an attacker could achieve)
- **Affected versions** (if known)
- **Suggested fix** (if you have one)
- **Your contact information** for follow-up questions

### What to Expect

**Initial Response**: Support acknowledges every report within **1 business day**, matching our standard support SLA.

**Assessment**: If your report describes a concrete, exploitable vulnerability, our security team provides an initial assessment within **1 week** of escalation, including:
- Confirmation that we can reproduce the issue
- Severity classification (Critical, High, Medium, Low)
- Expected timeline for a fix

**Fix Timeline**:
- **Critical vulnerabilities**: Days to 1 week
- **High vulnerabilities**: 1-2 weeks
- **Medium vulnerabilities**: 2-4 weeks
- **Low vulnerabilities**: Next regular release cycle

### Coordinated Disclosure

We work with you to coordinate public disclosure timing. We prefer to:
1. Develop and test a fix
2. Release the patch
3. Publish a security advisory
4. Credit you (if wanted) in the advisory

If you would like to be credited for the discovery, please let us know in your report. We acknowledge your contribution in:
- The GitHub Security Advisory
- The release notes

### Security Updates

Security fixes are released as:
- **Patch versions** (for example, v1.4.0 → v1.4.1)
- **GitHub Security Advisories** at https://github.com/SolaceProducts/terraform-provider-solacebroker/security/advisories
- The **Security** section of the GitHub release notes

Subscribe to releases on GitHub to be notified of security updates.

## Vulnerability Disclosure Policy

We follow **coordinated disclosure**:

1. **Private reporting period**: Security researchers report issues privately
2. **Fix development**: We develop and test patches without public disclosure
3. **Coordinated release**: We coordinate disclosure timing with the reporter
4. **Public disclosure**: After the fix is released, we publish a security advisory

We will not pursue legal action against security researchers who:
- Report vulnerabilities responsibly and privately
- Give us reasonable time to fix issues before public disclosure
- Do not exploit vulnerabilities beyond proof-of-concept
- Do not access, modify, or delete data belonging to others

## Questions?

If you have questions about this security policy, contact [support@solace.com](mailto:support@solace.com).

For non-security bugs and feature requests, please use [GitHub Issues](https://github.com/SolaceProducts/terraform-provider-solacebroker/issues).
