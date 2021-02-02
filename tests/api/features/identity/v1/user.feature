@identity
@user
Feature: get version
  @user-info
  Scenario: Get user info
    When I send "GET" request to "/identity/v1/user/1"
    Then the response status code should be 200
    And the response should match json:
      """
      {
        "status": 1,
        "code": "001.000.000",
        "message": "Success.",
        "data": {
            "id": 1,
            "full_name": "",
            "username": "lichtv",
            "email": "example@imgo.com",
            "created_at": "2020-11-16T02:10:38+07:00",
            "modified_at": "2020-11-16T02:10:38+07:00",
            "status": 0
          }
      }
      """
  
  @user-info-outline
  Scenario Outline: Get user info outline
    When I send "GET" request to "/identity/v1/user/<id>"
    Then the response status code should be <status-code>
    And the response should match json:
      """
      {
        "status": 1,
        "code": "001.000.000",
        "message": "Success.",
        "data": {
            "id": <id>,
            "full_name": "",
            "username": <username>,
            "email": <email>,
            "created_at": "2020-11-16T02:10:38+07:00",
            "modified_at": "2020-11-16T02:10:38+07:00",
            "status": 0
          }
      }
      """
  Examples: 
    | id | status-code | username | email |
    | 1 | 200 | "lichtv" | "example@imgo.com" |
    | 2 | 200 | "lichtv_test" | "example_test@imgo.com" |