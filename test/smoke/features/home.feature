@smoke @homepage
Feature: Homepage

Scenario: Opens the home
When I send "GET" request to "/"
Then The response code should be 200
