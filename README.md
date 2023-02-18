# TakeOver-v1

### What is Subdomain Takeover?

Subdomain Takeover is a type of vulnerability that arises when a subdomain points to an external service that has been deleted or is no longer in use. Common examples of these external services include Github, Heroku, Gitlab, and Tumblr. In this scenario, an attacker can exploit this vulnerability if the original owner fails to remove the DNS entry that points to the deleted service, allowing the attacker to takeover the subdomain by adding a CNAME file containing the subdomain name. This type of vulnerability can have significant security implications and requires careful attention to prevent exploitation.

Here is the command that checks CNAME record of a subdomain.

$dig CNAME apt.shopify.com --> apt.shopify.com.s3-website-us-east-1.amazonaws.com.

### How Can Takeover script help bug bounty hunters?

Managing and securing large numbers of subdomains can be a challenging task for organizations. In order to effectively monitor subdomains for potential security risks, it is important to have a tool that can automate the process of checking CNAME records for each domain. The script in question takes a file name as input, and performs a series of actions to produce output that displays the CNAME record for each domain in the input file. This approach enables security professionals to easily manage and monitor a large number of subdomains, and can help to identify potential vulnerabilities more efficiently.

### How can I recognise if the subdomain is vulnerable to subdomain takeover?

When a service is deleted, it is important to analyze the fingerprints that may be left behind when the DNS entry remains in place. In some cases, a vulnerable subdomain may display an error message when visited by an attacker, such as "There isn't a Github Pages site here." By carefully examining these error messages, security professionals can gain valuable insight into potential vulnerabilities that may exist within a domain or subdomain, and take steps to mitigate these risks. This type of analysis is essential for identifying and addressing security issues that may arise from improperly configured or abandoned services.

![Alt text](/53D457E2-57EC-4217-89EA-9A6DA84E1ACE.png?raw=true "TakeOver-v1" )

Security researcher @edoverflow has listed all services and their fingerprints. For more detail visit https://github.com/EdOverflow/can-i-take-over-xyz

# Youtube

[![Takeover-v1](https://img.youtube.com/vi/vpCI8foxP00/0.jpg)](https://www.youtube.com/watch?v=vpCI8foxP00)
