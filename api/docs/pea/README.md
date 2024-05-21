

# Folder: AD Accounts Backend API


**Authentication**: `bearer`<table>
        <tr>
        <th>AddTo</th>
        <th>Key</th>
        <th>Value</th>
        </tr>
        <tr>
        <td>HEADERS</td>
        <td>Authorization</td>
        <td>`bearer innodev1234`</td>
        </tr>
        </table>




---
## [GET] Get User

**Method**: GET

**RequestURL**:  `<<ENDPOINT>>/getUser`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`username`</td>
    <td><code>int</code></td>
    <td>`1234567890133`</td>
    </tr>
    </table>
**Authentication Type**: `none` 
    







**ContentType**: ``


**Request Body**:





---
## [GET] Get List User

**Method**: GET

**RequestURL**:  `<<ENDPOINT>>/listAllUser`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`ad_is_enabled`</td>
    <td><code>bool</code></td>
    <td>`true`</td>
    </tr>
    <tr>
    <td>`ad_locked`</td>
    <td><code>bool</code></td>
    <td>`true`</td>
    </tr>
    </table>
**Authentication Type**: `none` 
    







**ContentType**: ``


**Request Body**:





---
## [POST] Create User

**Method**: POST

**RequestURL**:  `<<ENDPOINT>>/createUser`


**Authentication Type**: `none` 
    







**ContentType**: `application/json`


**Request Body**:
```json
{
    "username": "9900000000001",
    "firstname": "ทดสอบ",
    "lastname": "ลืมรหัสผ่าน",
    "password": "12345678",
    "email": "99testforgotbgm2024@mailinator.com"
}
```






---
## [POST] Update User Info

**Method**: POST

**RequestURL**:  `<<ENDPOINT>>/updateUser`


**Authentication Type**: `none` 
    







**ContentType**: `application/json`


**Request Body**:
```json
{
    "username": "teeradet.huag",
    "firstname": "teeradet",
    "lastname": "phondatparinya",
    "email": "tee@mailinator.com"
}

```






---
## [POST] Reset Password

**Method**: POST

**RequestURL**:  `<<ENDPOINT>>/resetPassword`


**Authentication Type**: `none` 
    







**ContentType**: `application/json`


**Request Body**:
```json
{
    "username": "chakree.kae",
    "new_password": "12345"
}
```






---
## [POST] Change Password

**Method**: POST

**RequestURL**:  `<<ENDPOINT>>/changePassword`


**Authentication Type**: `none` 
    







**ContentType**: `application/json`


**Request Body**:
```json
{
    "username": "chakree.kae",
    "old_password": "123456",
    "new_password": "1234567"
}
```






---
## [GET] List Locked Out Users

**Method**: GET

**RequestURL**:  `<<ENDPOINT>>/listLockedUser`


**Authentication Type**: `none` 
    







**ContentType**: ``


**Request Body**:





---
## [POST] Unlock User

**Method**: POST

**RequestURL**:  `<<ENDPOINT>>/unlockUser`


**Authentication Type**: `none` 
    







**ContentType**: `application/json`


**Request Body**:
```json
{
    "username": "chakree.kae"
}
```






---
## [POST] Lock User

**Method**: POST

**RequestURL**:  `<<ENDPOINT>>/unlockUser`


**Authentication Type**: `none` 
    







**ContentType**: `application/json`


**Request Body**:
```json
{
    "username": "chakree.kae"
}
```






---
## [POST] Disable User

**Method**: POST

**RequestURL**:  `<<ENDPOINT>>/disableUser`


**Authentication Type**: `none` 
    







**ContentType**: `application/json`


**Request Body**:
```json
{
    "username": "chakree.kae"
}
```






---
## [POST] Enable User

**Method**: POST

**RequestURL**:  `<<ENDPOINT>>/enableUser`


**Authentication Type**: `none` 
    







**ContentType**: `application/json`


**Request Body**:
```json
{
    "username": "teeradet.huag"
}
```






---
## [DELETE] Delete User

**Method**: DELETE

**RequestURL**:  `<<ENDPOINT>>/deleteUser`


**Authentication Type**: `none` 
    







**ContentType**: `application/json`


**Request Body**:
```json
{
    "username": "8568778590095"
}
```






---
## [POST] 2Factor Enable

**Method**: POST

**RequestURL**:  `<<ENDPOINT>>/twoFactor_enable`


**Authentication Type**: `none` 
    







**ContentType**: `application/json`


**Request Body**:
```json
{
    "username": "teeradet.huag"
}
```






---
## [POST] 2Factor Disable

**Method**: POST

**RequestURL**:  `<<ENDPOINT>>/twoFactor_disable`


**Authentication Type**: `none` 
    







**ContentType**: `application/json`


**Request Body**:
```json
{
    "username": "teeradet.huag"
}
```






---
## [GET] 2Factor Get Secret (JSON)

**Method**: GET

**RequestURL**:  `<<ENDPOINT>>/twoFactor_get`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`username`</td>
    <td><code>string</code></td>
    <td>`teeradet.huag`</td>
    </tr>
    <tr>
    <td>`type`</td>
    <td><code>string</code></td>
    <td>`json`</td>
    </tr>
    </table>
**Authentication Type**: `none` 
    







**ContentType**: ``


**Request Body**:





---
## [GET] 2Factor Get Secret (QRCode)

**Method**: GET

**RequestURL**:  `<<ENDPOINT>>/twoFactor_get`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`username`</td>
    <td><code>string</code></td>
    <td>`teeradet.huag`</td>
    </tr>
    <tr>
    <td>`type`</td>
    <td><code>string</code></td>
    <td>`qrcode`</td>
    </tr>
    </table>
**Authentication Type**: `none` 
    







**ContentType**: ``


**Request Body**:





---
## [POST] 2Factor Test Check

**Method**: POST

**RequestURL**:  `<<ENDPOINT>>/twoFactor_check`


**Authentication Type**: `none` 
    







**ContentType**: `application/json`


**Request Body**:
```json
{
    "username": "chakree.kae",
    "token": 969019
}
```






---
## [GET] New Request

**Method**: GET

**RequestURL**:  ``


**Authentication Type**: `none` 
    







**ContentType**: ``


**Request Body**:





---
## [GET] New Request

**Method**: GET

**RequestURL**:  ``


**Authentication Type**: `none` 
    







**ContentType**: ``


**Request Body**:





---





# Folder: API


**Authentication**: `api-key`<table>
        <tr>
        <th>AddTo</th>
        <th>Key</th>
        <th>Value</th>
        </tr>
        <tr>
        <td>HEADERS</td>
        <td>`<<API_KEY_NAME>>`</td>
        <td>`<<API_KEY_VALUE>>`</td>
        </tr>
        </table>**Headers**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Value</th>
    </tr><tr>
    <td>`X-Test`</td>
    <td>`123`</td>
    </tr></table>




---





---

## Folder: AUTH
**Authentication**: `inherit`<p></p>






---

### [POST] /v1/signin

**Method**: POST

**RequestURL**: `<<BGM_AUTH>>/v1/signin`


**Authentication Type**: `inherit` 
    
**ContentType**: `application/json`

**Request Body**:
```json
{
    "provider": "KEYCLOAK",
    "host": "localhost"
}

```







---

### [GET] /v1/callback

**Method**: GET

**RequestURL**: `<<BGM_AUTH>>/v1/callback`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`state`</td>
    <td><code>unknow</code></td>
    <td> </td>
    </tr>
    <tr>
    <td>`session_state`</td>
    <td><code>unknow</code></td>
    <td> </td>
    </tr>
    <tr>
    <td>`code`</td>
    <td><code>unknow</code></td>
    <td> </td>
    </tr>
    </table>
**Authentication Type**: `none` 
    
**ContentType**: ``

**Request Body**:






---

### [POST] /v1/sign/token

**Method**: POST

**RequestURL**: `<<BGM_AUTH>>/v1/sign/token`


**Authentication Type**: `none` 
    
**ContentType**: `application/json`

**Request Body**:
```json
{
    "provider_token": "eyJhbGciOiJIUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJkMjcyMzhkZi1jNDdlLTQ4NjktOGNhOS1kODM3ZDFhYWM5OTMifQ.eyJleHAiOjE3MTA5MjMxMDYsImlhdCI6MTcxMDkyMTMwNiwianRpIjoiOWExNDZkMDktNzg5OS00MTkxLWI5OTQtNmViNDBlODYxZjAwIiwiaXNzIjoiaHR0cHM6Ly9iZ20tc3NvLmlubm92YXNpdmUuZGV2L3JlYWxtcy9CR01TU08iLCJhdWQiOiJodHRwczovL2JnbS1zc28uaW5ub3Zhc2l2ZS5kZXYvcmVhbG1zL0JHTVNTTyIsInN1YiI6IjI3NTViOWNlLWRlZjktNDdlYi1iNDUxLWM2YTljNjY1YzExMSIsInR5cCI6IlJlZnJlc2giLCJhenAiOiJiZ20tc2VydmljZSIsInNlc3Npb25fc3RhdGUiOiJmZjcxM2FkYi04YTFhLTQ3OTctODI0Mi0zODhiOTE4NDZjZTAiLCJzY29wZSI6Im9wZW5pZCBlbWFpbCBtaWNyb3Byb2ZpbGUtand0IHByb2ZpbGUiLCJzaWQiOiJmZjcxM2FkYi04YTFhLTQ3OTctODI0Mi0zODhiOTE4NDZjZTAifQ.Z0zatHvVfkZjSnahJNuQFfxbF5cNZZqoGDlBlhsmcss",
    "sso_client_id": "8ef5b744-0b33-4768-bf8c-d082c39cf533"
}

```







---

### [POST] /v1/signout

**Method**: POST

**RequestURL**: `<<BGM_AUTH>>/v1/signout`


**Authentication Type**: `bearer` 
    

**BearerToken**: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7Imp0aSI6IjBjMTdmODRhLWNmYzEtNGExZC1iNWIzLWE3ZTAyZjdkNWJiZSIsImlkIjoiZmUyZTI4MDEtNGU1Ni00OTQ3LTgwMzQtNDBiMWIwMDEyZTJjIiwidXNlcm5hbWUiOiIyMDgxNTA1MDAyNjQ0Iiwic3NvX2NsaWVudF9pZCI6IjA5OGY4MzlkLTRmM2ItNDBkMy04NjczLWIzOThiMzcyZTBkMSJ9LCJjbGFpbXMiOnsiaXNzIjoiYmdtLWFwaS1hdXRoIiwiZXhwIjoxNzExNTIzNTI0LCJpYXQiOjE3MTE0MzcxMjQsImp0aSI6IjBjMTdmODRhLWNmYzEtNGExZC1iNWIzLWE3ZTAyZjdkNWJiZSJ9fQ.ljumV8JTeaLF6KHZRw5I4fxaHHuRPShiKk089F7F2YY`

**ContentType**: ``

**Request Body**:






---

### [POST] /v1/resign/token

**Method**: POST

**RequestURL**: `<<BGM_AUTH>>/v1/resign/token`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:






---

### [GET] /v1/roles/config

**Method**: GET

**RequestURL**: `http://127.0.0.1:3001/v1/roles/config`


**Authentication Type**: `bearer` 
    
**ContentType**: ``

**Request Body**:






---

### [POST] /v1/create/user/role

**Method**: POST

**RequestURL**: `<<BGM_AUTH>>/v1/create/user/role`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: `application/json`

**Request Body**:
```json
{
    "email": "testcreatepermission@mailinator.com",
    "role_ids": [
        "e2e6531a-d654-4cdc-910f-81f8d4443e1e",
        "605bd7a8-709a-4a44-adb3-9319a6bbbfd2"
    ]
}

```







---

### [PUT] /v1/edit/user/:user_id/role

**Method**: PUT

**RequestURL**: `<<BGM_AUTH>>/v1/edit/user/:user_id/role`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: `application/json`
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>user_id</code></td>
    <td><code>aef5c954-ede8-4bc3-b0a6-1c93a502b206</code></td>
    </tr>
</table>

**Request Body**:
```json
{
    "role_ids": ["605bd7a8-709a-4a44-adb3-9319a6bbbfd2"]
}
```







---

### [PUT] /v1/status/user/:user_id/role

**Method**: PUT

**RequestURL**: `<<BGM_AUTH>>/v1/status/user/:user_id/role`


**Authentication Type**: `none` 
    
**ContentType**: `application/json`
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>user_id</code></td>
    <td><code>991697b8-1cf6-424f-9dec-d2bf0d2f93af</code></td>
    </tr>
</table>

**Request Body**:
```json
{
    "status": "ACTIVE" // ACTIVE || INACTIVE
}
```







---

### [DELETE] /v1/delete/user/:user_id/role

**Method**: DELETE

**RequestURL**: `<<BGM_AUTH>>/v1/delete/user/:user_id/role`


**Authentication Type**: `none` 
    
**ContentType**: `application/json`
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>user_id</code></td>
    <td><code></code></td>
    </tr>
</table>

**Request Body**:
```json
```







---

## Folder: ORGANIZATION
**Authentication**: `inherit`<p></p>






---

### [GET] /v1/organizes

**Method**: GET

**RequestURL**: `<<BGM_ORGANIZATION>>/v1/organizes`


**Headers**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>X-Auth</td>
    <td>`hwz9a8wumM6G`</td>
    </tr></table>
**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`per_page`</td>
    <td><code>int</code></td>
    <td>`20`</td>
    </tr>
    <tr>
    <td>`org_type`</td>
    <td><code>string</code></td>
    <td>`FACULTY`</td>
    </tr>
    <tr>
    <td>`page`</td>
    <td><code>int</code></td>
    <td>`1`</td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7Imp0aSI6IjhlY2JjMjU4LTYyOGMtNGZhZi05NTNkLWQwMzUyNmM5ZjVhMiIsImlkIjoiZjA2NzExZWQtNzcwYS00NzU0LWEyYTItN2JjM2FiZGE2MTY0IiwidXNlcm5hbWUiOiJuZXd0ZXN0MTAxQGttdXR0LmFjLnRoIiwic3NvX2NsaWVudF9pZCI6ImZkZDMzOTQxLThhOGEtNGNlMS1hYzlkLWI0YmE4NjU3MzZlYyJ9LCJjbGFpbXMiOnsiaXNzIjoiYmdtLWFwaS1hdXRoIiwiZXhwIjoxNzE1NzYzMDI0LCJpYXQiOjE3MTU2NzY2MjQsImp0aSI6IjhlY2JjMjU4LTYyOGMtNGZhZi05NTNkLWQwMzUyNmM5ZjVhMiJ9fQ.GpU1XsOn6RLt2TL7A2jmVs-7sz25eDIfRgLUXs17zu0`

**ContentType**: ``

**Request Body**:






---

### [GET] /v1/organize/:organize_id

**Method**: GET

**RequestURL**: `<<BGM_ORGANIZATION>>/v1/organize/:organize_id`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>organize_id</code></td>
    <td><code>ed6fd893-13e5-4699-bee7-80ac119fc79a</code></td>
    </tr>
</table>

**Request Body**:






---

### [GET] /v1/organize/:organize_id/departments

**Method**: GET

**RequestURL**: `<<BGM_ORGANIZATION>>/v1/organize/:organize_id/departments`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`campus_code`</td>
    <td><code>int</code></td>
    <td>`01`</td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>organize_id</code></td>
    <td><code>f6201566-a552-4ba4-v-3990e7715845</code></td>
    </tr>
</table>

**Request Body**:






---

### [GET] /v1/buildings

**Method**: GET

**RequestURL**: `<<BGM_ORGANIZATION>>/v1/buildings`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`building_id`</td>
    <td><code>uuid</code></td>
    <td>`9eca30a9-0514-4d6b-8046-515dfe463bf3`</td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:






---

## Folder: USER
**Authentication**: `inherit`<p></p>



---

### Folder: consent
**Authentication**: `inherit`<p></p>

---

#### [POST] /v1/management/consent

**Method**: POST

**RequestURL**: `<<BGM_USER>>/v1/management/consent`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: `application/json`

**Request Body**:

```json

{
    "detail": "ทดสอบ-3"
}

```



---

#### [GET] /v1/management/consent

**Method**: GET

**RequestURL**: `<<BGM_USER>>/v1/management/consent`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`version`</td>
    <td><code>unknow</code></td>
    <td> </td>
    </tr>
    <tr>
    <td>`page`</td>
    <td><code>int</code></td>
    <td>`1`</td>
    </tr>
    <tr>
    <td>`per_page`</td>
    <td><code>int</code></td>
    <td>`10`</td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:



---

#### [GET] /v1/management/:consent_id/consent

**Method**: GET

**RequestURL**: `<<BGM_USER>>/v1/management/:consent_id/consent`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: `application/json`
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>consent_id</code></td>
    <td><code>cac3a4f7-9fcd-42ef-97af-3874ef7c5758</code></td>
    </tr>
</table>

**Request Body**:

```json

```



---

#### [PUT] /v1/management/status/:consent_id/consent

**Method**: PUT

**RequestURL**: `<<BGM_USER>>/v1/management/status/:consent_id/consent`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: `application/json`
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>consent_id</code></td>
    <td><code>cac3a4f7-9fcd-42ef-97af-3874ef7c5758</code></td>
    </tr>
</table>

**Request Body**:

```json

{
    "status": "ON" // ON || OFF
}

```



---

#### [GET] /v1/consent

**Method**: GET

**RequestURL**: `<<BGM_USER>>/v1/consent`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`consent_type`</td>
    <td><code>unknow</code></td>
    <td> </td>
    </tr>
    </table>
**Authentication Type**: `none` 
    
**ContentType**: ``

**Request Body**:



---

#### [PUT] /v1/consent

**Method**: PUT

**RequestURL**: `<<BGM_USER>>/v1/consent`


**Authentication Type**: `none` 
    
**ContentType**: `application/json`

**Request Body**:

```json

{
    "consent_id": "uuid"
}

```



---

#### [GET] /v1/my/consent

**Method**: GET

**RequestURL**: `<<BGM_USER>>/v1/my/consent`


**Authentication Type**: `none` 
    
**ContentType**: ``

**Request Body**:








---

### [GET] /v1/me

**Method**: GET

**RequestURL**: `<<BGM_USER>>/v1/me`


**Authentication Type**: `bearer` 
    

**BearerToken**: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7Imp0aSI6IjcyN2NiODRkLTA5ZGUtNDcyYy1iMWZjLTY4MTc0ZTQ2YTQyMCIsImlkIjoiZDA4NDIzMTItZGFjMy00MjY0LWFkYjEtNDJmNDEyNzJiMmNhIiwidXNlcm5hbWUiOiJ0ZWVyYWRldC5odWFnLm1haWxAa211dHQuYWMudGgifSwiY2xhaW1zIjp7ImlzcyI6ImJnbS1hcGktYXV0aCIsImV4cCI6MTcxNjE4Nzc5OCwiaWF0IjoxNzEwMTM5Nzk4LCJqdGkiOiI3MjdjYjg0ZC0wOWRlLTQ3MmMtYjFmYy02ODE3NGU0NmE0MjAifX0.D1BWx5E606mAuvXQ3HHksWyl5sq9LTH6wDmtwzyCThQ`

**ContentType**: ``

**Request Body**:






---

### [GET] /v1/types

**Method**: GET

**RequestURL**: `<<BGM_USER>>/v1/types`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`auth_type`</td>
    <td><code>string</code></td>
    <td>`KMUTT`</td>
    </tr>
    <tr>
    <td>`code`</td>
    <td><code>string</code></td>
    <td>`STUDENT`</td>
    </tr>
    <tr>
    <td>`search_word`</td>
    <td><code>string</code></td>
    <td>`สัญญา`</td>
    </tr>
    <tr>
    <td>`page`</td>
    <td><code>int</code></td>
    <td>`1`</td>
    </tr>
    <tr>
    <td>`per_page`</td>
    <td><code>int</code></td>
    <td>`10`</td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7Imp0aSI6ImVmYjNhZDc4LWJmYzYtNGM5OS04OGUxLTA2ZjM0ZGU5Y2RjYiIsImlkIjoiZjA2NzExZWQtNzcwYS00NzU0LWEyYTItN2JjM2FiZGE2MTY0IiwidXNlcm5hbWUiOiJuZXd0ZXN0MTAxQGttdXR0LmFjLnRoIiwic3NvX2NsaWVudF9pZCI6IjE0Y2JmYjQ1LTAxY2ItNDhjZi1iY2MyLWUyZTFkNzUzMTljYSJ9LCJjbGFpbXMiOnsiaXNzIjoiYmdtLWFwaS1hdXRoIiwiZXhwIjoxNzE0Nzk0OTAwLCJpYXQiOjE3MTQ3MDg1MDAsImp0aSI6ImVmYjNhZDc4LWJmYzYtNGM5OS04OGUxLTA2ZjM0ZGU5Y2RjYiJ9fQ.h6d6xi-svMAI2V-SYV0o-fLPrF3Kdg8Ik80C1s1C7Jo`

**ContentType**: ``

**Request Body**:






---

### [GET] /v1/lists

**Method**: GET

**RequestURL**: `<<BGM_USER>>/v1/lists`


**Headers**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>X-Auth</td>
    <td>`hwz9a8wumM6G`</td>
    </tr></table>
**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`search_word`</td>
    <td><code>string</code></td>
    <td>`Jame`</td>
    </tr>
    <tr>
    <td>`user_type_id`</td>
    <td><code>uuid</code></td>
    <td>`88c8a7f6-2fe0-459c-a54d-0fd8564a1f0e`</td>
    </tr>
    <tr>
    <td>`auth_type`</td>
    <td><code>string</code></td>
    <td>`KMUTT`</td>
    </tr>
    <tr>
    <td>`page`</td>
    <td><code>int</code></td>
    <td>`1`</td>
    </tr>
    <tr>
    <td>`per_page`</td>
    <td><code>int</code></td>
    <td>`20`</td>
    </tr>
    <tr>
    <td>`org_id`</td>
    <td><code>uuid</code></td>
    <td>`fedc4633-d754-435d-a821-f78ee3d49592`</td>
    </tr>
    <tr>
    <td>`renewable_pink_orange`</td>
    <td><code>bool</code></td>
    <td>`true`</td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7Imp0aSI6IjI5ZTVjNDliLWIwZDgtNDgxMy05YTQ4LTJkM2JjODg0Zjg3OCIsImlkIjoiZjA2NzExZWQtNzcwYS00NzU0LWEyYTItN2JjM2FiZGE2MTY0IiwidXNlcm5hbWUiOiJuZXd0ZXN0MTAxQGttdXR0LmFjLnRoIiwic3NvX2NsaWVudF9pZCI6ImYwODk0NDhlLTY4YzMtNDgyNy04NzY2LWZlNDg5ZjkzNDE5ZiJ9LCJjbGFpbXMiOnsiaXNzIjoiYmdtLWFwaS1hdXRoIiwiZXhwIjoxNzE1OTI3MTQ1LCJpYXQiOjE3MTU4NDA3NDUsImp0aSI6IjI5ZTVjNDliLWIwZDgtNDgxMy05YTQ4LTJkM2JjODg0Zjg3OCJ9fQ.ab9aETneoMm4Ux6rJ1g1UdBQ0kYON-gbg3vm164yKYY`

**ContentType**: `application/json`

**Request Body**:
```json
```







---

### [GET] /v1/:user_id/details

**Method**: GET

**RequestURL**: `<<BGM_USER>>/v1/:user_id/details`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>user_id</code></td>
    <td><code>388decb0-7582-44a6-9923-77640c5b063f</code></td>
    </tr>
</table>

**Request Body**:






---

### [POST] /v1/accounts

**Method**: POST

**RequestURL**: `<<BGM_USER>>/v1/accounts`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: `application/json`

**Request Body**:
```json
{
    "organize_id": "cc86b6ed-0db4-41cb-a9d2-60e73b13abd2",
    "fullname_th": "นายทดสอบ sun",
    "citizen": "1000000000002",
    "contact_email_1": "sun_test_2@mailinator.com",
    "code": "9900002",
    "user_type_id": "1f3c95f8-fde1-4589-8e69-5cf85f904ade",
    "department_id": "5e2069ef-a154-44be-bd50-98c4c924a2cd",
    "contract": null,
    "cover": null,
    "expired_at": "2024-10-31 23:59:00"
}
```







---

### [PUT] /v1/:user_id/accounts

**Method**: PUT

**RequestURL**: `<<BGM_USER>>/v1/:user_id/accounts`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: `application/json`
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>user_id</code></td>
    <td><code>8353791e-3b5b-47cd-9e47-2cec04ed576c</code></td>
    </tr>
</table>

**Request Body**:
```json
{
    "organize_id": "c7345793-7980-4e76-8db6-abb0ccbe3719",                  // หน่วยงาน
    "fullname_th": "ปรียาภัทร นะยอกแล้ว",                                       // ชื่อนามสกุล
    "citizen": "8568778590039",                                             // หมายเลขบัตรประชาชน
    "contact_email_1": "teeradet.huag@gmail.com",                         // อีเมล์ผู้ใช้
    "code": "6412345678",                                                   // รหัสพนักงาน/รหัสนักศึกษา
    "user_type_id": "19e7a6ab-631e-41a9-b0fb-36c54acdac62",                 // ประเภทผู้ใช้งาน
    "department_id": "2eb34215-aff1-461e-a3cf-b5b3bed3c685",                // หน่วยงานที่สังกัด
    "expired_at": "2024-03-05 23:59:00",                                     // วันเวลาที่หมดอายุ
    "contract": {
        "url": "https://bgm-s3.innovasive.dev/raw/contract/20240115_a0259c72-162f-44bd-9195-ccd3883c4de8_946263948.pdf?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=app%2F20240115%2Fap-southeast-1%2Fs3%2Faws4_request&X-Amz-Date=20240115T094245Z&X-Amz-Expires=10800&X-Amz-SignedHeaders=host&X-Amz-Signature=25ba65755f676560df06af4197cb02d33b79779ec811baa3614fde125a68987d",
        "content_type": "application/pdf",
        "size": 3028,
        "original_filename": "sample.pdf",
        "note": ""
    },
    "cover": null
}
```







---

### [POST] /v1/:user_id/accounts

**Method**: POST

**RequestURL**: `<<BGM_USER>>/v1/accounts`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: `application/json`

**Request Body**:
```json
{
    "organize_id": "c7345793-7980-4e76-8db6-abb0ccbe3719",                  // หน่วยงาน
    "fullname_th": "ปรียาภัทร เหล่าอรุณ",                                       // ชื่อนามสกุล
    "citizen": "8568778590095",                                             // หมายเลขบัตรประชาชน
    "contact_email_1": "Preyapat.l@mailinator.com",                         // อีเมล์ผู้ใช้
    "code": "6412345678",                                                   // รหัสพนักงาน/รหัสนักศึกษา
    "user_type_id": "19e7a6ab-631e-41a9-b0fb-36c54acdac62",                 // ประเภทผู้ใช้งาน
    "department_id": "2eb34215-aff1-461e-a3cf-b5b3bed3c685",                // หน่วยงานที่สังกัด
    "expired_at": "2024-12-05 23:59:00",                                     // วันเวลาที่หมดอายุ
    "contract": {
        "url": "https://bgm-s3.innovasive.dev/raw/contract/20240115_a0259c72-162f-44bd-9195-ccd3883c4de8_946263948.pdf?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=app%2F20240115%2Fap-southeast-1%2Fs3%2Faws4_request&X-Amz-Date=20240115T094245Z&X-Amz-Expires=10800&X-Amz-SignedHeaders=host&X-Amz-Signature=25ba65755f676560df06af4197cb02d33b79779ec811baa3614fde125a68987d",
        "content_type": "application/pdf",
        "size": 3028,
        "original_filename": "sample.pdf",
        "note": ""
    },
    "cover": {
        "url": "https://bgm-s3.innovasive.dev/raw/cover/20240115_703f468e-1b61-422a-b5cf-3aaa4c85ce4f_691303704.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=app%2F20240115%2Fap-southeast-1%2Fs3%2Faws4_request&X-Amz-Date=20240115T095152Z&X-Amz-Expires=10800&X-Amz-SignedHeaders=host&X-Amz-Signature=b4a66759c9886bcbda925739c0e433712485e541037e6242613032dd5d05aafc",
        "content_type": "image/jpeg",
        "size": 343694,
        "original_filename": "pikachu.jpeg",
        "note": ""
    }
}
```







---

### [POST] /v1/account/reset/passwords

**Method**: POST

**RequestURL**: `<<BGM_USER>>/v1/account/reset/passwords`


**Authentication Type**: `none` 
    
**ContentType**: `application/json`

**Request Body**:
```json
{
    "mail_code": "a00cb431-eab3-4e67-bd3a-7e6056cb6844",     // verified email id ส่งมาพร้อมกับ query params
    "password": "Az123456789#",
    "confirm_password": "Az123456789#"
}
```







---

### [POST] /v1/:user_id/account/request/reset/passwords

**Method**: POST

**RequestURL**: `<<BGM_USER>>/v1/:user_id/account/request/reset/passwords`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: `application/json`
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>user_id</code></td>
    <td><code>ba10065a-bb6b-4742-9fb5-e7e9a16d2a0d</code></td>
    </tr>
</table>

**Request Body**:
```json
{
    "password": "Inno_4321",
    "confirm_password": "Inno_4321",
    "updated_by": "ADMIN"
}
```







---

### [DELETE] /v1/:user_id

**Method**: DELETE

**RequestURL**: `<<BGM_USER>>/v1/:user_id`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>user_id</code></td>
    <td><code>4ca603b4-3d59-4275-87f0-576145d0ba67</code></td>
    </tr>
</table>

**Request Body**:






---

### [POST] /v1/account/management

**Method**: POST

**RequestURL**: `<<BGM_USER>>/v1/account/management`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: `application/x-www-form-urlencoded`

**Request Body**:<table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
        <tr>
        <td><code>data</code></td>
        <td><code>object</code></td>
        <td><code>{"email" : "Preyapat.l@mailinator.com","role_ids": ["782287b4-2eb8-4f56-a962-e321ac62adab","26bdd3f6-d0b7-48ae-b5a1-746b873bc959"]}</code></td>
        </tr>
        <tr>
        <td><code>hello</code></td>
        <td><code>string</code></td>
        <td><code>worl</code></td>
        </tr>
        <tr>
        <td><code>foo</code></td>
        <td><code>string</code></td>
        <td><code>ba</code></td>
        </tr>
        <tr>
        <td><code>abc_123</code></td>
        <td><code>int</code></td>
        <td><code>45</code></td>
        </tr>
    </table><!-- tabs:start -->
**Example Response**#### **200 SUCCESS**
**Response Header** 


**Response Body**
```json
{"message": "ok"}
```#### **409 USER DUPLICATE**
**Response Header** 


**Response Body**
```json
{"message": "user_duplicate"}
```#### **403 FORBIDDEN**
**Response Header** 


**Response Body**
```json
{"message": "forbidden"}
```<!-- tabs:end -->






---

### [POST] /v1/account/request/reset/passwords

**Method**: POST

**RequestURL**: `<<BGM_USER>>/v1/account/request/reset/passwords`


**Authentication Type**: `none` 
    
**ContentType**: `application/json`

**Request Body**:
```json
{
    "username": "1103100478417" // citizen_id
}
```







---

### [POST] /v1/reset/password/check/ref_code

**Method**: POST

**RequestURL**: `<<BGM_USER>>/v1/reset/password/check/ref_code`


**Authentication Type**: `none` 
    
**ContentType**: `application/json`

**Request Body**:
```json
{
    "code": "",
    "ref_code": "",
    "username": "9900000000001" // citizen_id
}
```







---

### [POST] /v1/reset/password

**Method**: POST

**RequestURL**: `<<BGM_USER>>/v1/reset/password`


**Authentication Type**: `none` 
    
**ContentType**: `application/json`

**Request Body**:
```json
{
    "username": "1000000000001", //citizen_id
    "password": "P@ssw0rd3",
    "confirm_password": "P@ssw0rd3",
    "code": "835573",
    "ref_code": "DY60"
}
```







---

### [GET] /v1/list/detail

**Method**: GET

**RequestURL**: `<<BGM_USER>>/v1/lists/detail`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`user_id`</td>
    <td><code>uuid</code></td>
    <td>`d07ceb76-d40e-4fcb-9d2c-6dbf47dd389b`</td>
    </tr>
    <tr>
    <td>`user_id`</td>
    <td><code>uuid</code></td>
    <td>`de284c5b-4841-4c07-88cd-44f6aa0bbeda`</td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:






---

### [POST] /v1/users/file

**Method**: POST

**RequestURL**: `<<BGM_USER>>/v1/users/file`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`mode`</td>
    <td><code>string</code></td>
    <td>`renew`</td>
    </tr>
    <tr>
    <td>`mode`</td>
    <td><code>string</code></td>
    <td>`notification_recipient`</td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: `multipart/form-data`

**Request Body**:<table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
        <tr>
        <td><code>file</code></td>
        <td><code> unknow</code></td>
        <td><code></code></td>
        </tr>
        <tr>
        <td><code>data</code></td>
        <td><code> object</code></td>
        <td><code>{"module": "user", "key": "user_list"}</code></td>
        </tr>
    </table>






---

### [GET] /v1/users/role/list

**Method**: GET

**RequestURL**: `<<BGM_USER>>/v1/users/role/list`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`search_world`</td>
    <td><code>unknow</code></td>
    <td> </td>
    </tr>
    <tr>
    <td>`role_id`</td>
    <td><code>unknow</code></td>
    <td> </td>
    </tr>
    <tr>
    <td>`page`</td>
    <td><code>int</code></td>
    <td>`1`</td>
    </tr>
    <tr>
    <td>`per_page`</td>
    <td><code>int</code></td>
    <td>`10`</td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:






---

### [GET] New Request

**Method**: GET

**RequestURL**: ``


**Authentication Type**: `none` 
    
**ContentType**: ``

**Request Body**:






---

## Folder: FILE
**Authentication**: `inherit`<p></p>






---

### [POST] /v1/uploads

**Method**: POST

**RequestURL**: `<<BGM_FILE>>/v1/uploads`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: `multipart/form-data`
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>MOCK_200_RESP</code></td>
    <td><code>{"a": "b"}</code></td>
    </tr>
</table>

**Request Body**:<table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
        <tr>
        <td><code>data</code></td>
        <td><code> object</code></td>
        <td><code>{"module": "announcements", "key": "cover"}</code></td>
        </tr>
        <tr>
        <td><code>file</code></td>
        <td><code> @file </code></td>
        <td><code>Web Browser File</code></td>
        </tr>
    </table>






---

### [POST] /v1/attach/uploads

**Method**: POST

**RequestURL**: `<<BGM_FILE>>/v1/attach/uploads`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: `multipart/form-data`

**Request Body**:<table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
        <tr>
        <td><code>data</code></td>
        <td><code> object</code></td>
        <td><code>{"module": "announcements", "key": "attach"}</code></td>
        </tr>
        <tr>
        <td><code>file</code></td>
        <td><code> unknow</code></td>
        <td><code></code></td>
        </tr>
    </table>






---

### [POST] /v1/attach/delete

**Method**: POST

**RequestURL**: `<<BGM_FILE>>/v1/attach/delete`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: `application/json`

**Request Body**:
```json
{
    "url": "https://bgm-s3.innovasive.dev/announcements/attach/20240402_ceb89c00-52c5-40ac-97e1-ca8870a28bde_971995670.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=app%2F20240402%2Fap-southeast-1%2Fs3%2Faws4_request&X-Amz-Date=20240402T031540Z&X-Amz-Expires=10800&X-Amz-SignedHeaders=host&X-Amz-Signature=cdff55d44eb8fd08a3eb107519fe42703d284275c7ec3d7e2a39ac2e2bdb1a56"
}
```







---

## Folder: FACILITY
**Authentication**: `inherit`<p></p>






---

### [GET] /v1/global/configs

**Method**: GET

**RequestURL**: `<<BGM_FACILITY>>/v1/global/configs`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:






---

### [GET] /v1/parking/penalities

**Method**: GET

**RequestURL**: `<<BGM_FACILITY>>/v1/parking/penalities`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`code`</td>
    <td><code>string</code></td>
    <td>`red`</td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:






---

### [GET] /v1/parking/permit/types

**Method**: GET

**RequestURL**: `<<BGM_FACILITY>>/v1/parking/permit/types`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`user_type_id`</td>
    <td><code>uuid</code></td>
    <td>`6126f9e3-5c26-42a3-bb15-8a4455fbfb4a`</td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:






---

### [GET] /v1/parking/permit/:parking_id

**Method**: GET

**RequestURL**: `<<BGM_FACILITY>>/v1/parking/permit/:parking_id`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>parking_id</code></td>
    <td><code>6ef3be13-bcbd-4eed-ae53-dca2402494a9</code></td>
    </tr>
</table>

**Request Body**:






---

### [GET] /v1/app/parking/permit/:parking_id

**Method**: GET

**RequestURL**: `<<BGM_FACILITY>>/v1/app/parking/permit/:parking_id`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>parking_id</code></td>
    <td><code>cce1b229-0ad0-46ea-8b71-81cd307a9d3c</code></td>
    </tr>
</table>

**Request Body**:






---

### [GET] /v1/my/services ที่รวมใบอนุญาตของตัวเอง

**Method**: GET

**RequestURL**: `<<BGM_FACILITY>>/v1/my/services`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`statuses`</td>
    <td><code>string</code></td>
    <td>`APPROVED`</td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:






---

### [GET] /v1/parking/permits

**Method**: GET

**RequestURL**: `<<BGM_FACILITY>>/v1/parking/permits`


**Headers**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>X-Auth</td>
    <td>`hwz9a8wumM6G`</td>
    </tr></table>
**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`page`</td>
    <td><code>int</code></td>
    <td>`1`</td>
    </tr>
    <tr>
    <td>`per_page`</td>
    <td><code>int</code></td>
    <td>`20`</td>
    </tr>
    <tr>
    <td>`statuses`</td>
    <td><code>string</code></td>
    <td>`APPROVED`</td>
    </tr>
    <tr>
    <td>`statuses`</td>
    <td><code>string</code></td>
    <td>`REQUESTED`</td>
    </tr>
    <tr>
    <td>`is_expired`</td>
    <td><code>bool</code></td>
    <td>`false`</td>
    </tr>
    <tr>
    <td>`created_start_date`</td>
    <td><code>string</code></td>
    <td>`2024-02-06`</td>
    </tr>
    <tr>
    <td>`created_end_date`</td>
    <td><code>string</code></td>
    <td>`2024-02-06`</td>
    </tr>
    <tr>
    <td>`permit_type_codes`</td>
    <td><code>string</code></td>
    <td>`ORANGE`</td>
    </tr>
    <tr>
    <td>`permit_type_codes`</td>
    <td><code>string</code></td>
    <td>`PINK`</td>
    </tr>
    <tr>
    <td>`permit_type_codes`</td>
    <td><code>string</code></td>
    <td>`YELLOW`</td>
    </tr>
    <tr>
    <td>`statuses`</td>
    <td><code>string</code></td>
    <td>`REJECTED`</td>
    </tr>
    <tr>
    <td>`mode`</td>
    <td><code>string</code></td>
    <td>`BGM`</td>
    </tr>
    <tr>
    <td>`create_status`</td>
    <td><code>string</code></td>
    <td>`NEW`</td>
    </tr>
    <tr>
    <td>`user_type_id`</td>
    <td><code>uuid</code></td>
    <td>`92fe6591-7817-4c20-9c83-473798070cb5`</td>
    </tr>
    <tr>
    <td>`org_id`</td>
    <td><code>uuid</code></td>
    <td>`f53c0bee-c76f-4f15-939e-bbd9432bc36a`</td>
    </tr>
    <tr>
    <td>`search_word`</td>
    <td><code>string</code></td>
    <td>`ฟกอฟ`</td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7Imp0aSI6IjI5ZTVjNDliLWIwZDgtNDgxMy05YTQ4LTJkM2JjODg0Zjg3OCIsImlkIjoiZjA2NzExZWQtNzcwYS00NzU0LWEyYTItN2JjM2FiZGE2MTY0IiwidXNlcm5hbWUiOiJuZXd0ZXN0MTAxQGttdXR0LmFjLnRoIiwic3NvX2NsaWVudF9pZCI6ImYwODk0NDhlLTY4YzMtNDgyNy04NzY2LWZlNDg5ZjkzNDE5ZiJ9LCJjbGFpbXMiOnsiaXNzIjoiYmdtLWFwaS1hdXRoIiwiZXhwIjoxNzE1OTI3MTQ1LCJpYXQiOjE3MTU4NDA3NDUsImp0aSI6IjI5ZTVjNDliLWIwZDgtNDgxMy05YTQ4LTJkM2JjODg0Zjg3OCJ9fQ.ab9aETneoMm4Ux6rJ1g1UdBQ0kYON-gbg3vm164yKYY`

**ContentType**: ``

**Request Body**:






---

### [GET] /v1/parking/province

**Method**: GET

**RequestURL**: `<<BGM_FACILITY>>/v1/parking/province`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:






---

### [GET] /v1/parking/status

**Method**: GET

**RequestURL**: `<<BGM_FACILITY>>/v1/parking/permit/status`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:






---

### [GET] /v1/parking/fueltype

**Method**: GET

**RequestURL**: `<<BGM_FACILITY>>/v1/parking/fueltype`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:






---

### [GET] v1/parking/permit/:parking_id/car

**Method**: GET

**RequestURL**: `<<BGM_FACILITY>>/v1/parking/permit/:parking_id/car`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>parking_id</code></td>
    <td><code>b4cee74e-211b-43e3-8c48-43f3df12e321</code></td>
    </tr>
</table>

**Request Body**:






---

### [GET] v1/app/parking/permit/:parking_id/car

**Method**: GET

**RequestURL**: `<<BGM_FACILITY>>/v1/app/parking/permit/:parking_id/car`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>parking_id</code></td>
    <td><code>b4cee74e-211b-43e3-8c48-43f3df12e321</code></td>
    </tr>
</table>

**Request Body**:






---

### [POST] /v1/parking/permits

**Method**: POST

**RequestURL**: `<<BGM_FACILITY>>/v1/parking/permits`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: `application/json`

**Request Body**:
```json
{
    "user_type_id": "90f1e1a7-d929-4706-a345-f4a844114239",
    "organize_id": "c7345793-7980-4e76-8db6-abb0ccbe3719",
    "contact_number": "088-888-8888",
    "contact_email": "example@kmutt.ac.th",
    "payment_proof": {
        "url": "https://bgm-s3.innovasive.dev/raw/20240113_86fc3d35-36b2-4569-94d1-b1ccfd99c3f9_745240936.jpg",
        "content_type": "image/jpeg",
        "size": 10150,
        "note": ""
    }
}
```







---

### [POST] /v1/parking/permits/:parking_id/car

**Method**: POST

**RequestURL**: `<<BGM_FACILITY>>/v1/parking/permit/:parking_id/car`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`page`</td>
    <td><code>int</code></td>
    <td>`1`</td>
    </tr>
    <tr>
    <td>`per_page`</td>
    <td><code>int</code></td>
    <td>`10`</td>
    </tr>
    <tr>
    <td>`statuses`</td>
    <td><code>string</code></td>
    <td>`VERIFIED`</td>
    </tr>
    <tr>
    <td>`statuses`</td>
    <td><code>string</code></td>
    <td>`REQUESTED`</td>
    </tr>
    <tr>
    <td>`is_expired`</td>
    <td><code>bool</code></td>
    <td>`false`</td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: `application/json`
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>parking_id</code></td>
    <td><code>a32070d1-7ac6-4cd4-8c5f-088f7bf62cd7</code></td>
    </tr>
</table>

**Request Body**:
```json
{
    "cars" : [
        {
        "id": "2a3067cf-0794-4ca2-aca4-a9b506ce02ce",
        "brand": "b1",
        "model": "m1",
        "color": "c1",
        "year": "2567",
        "capacity": "1500",
        "license_number": "กก2asf023",
        "province": "กรุงเทพมหานคร",
        "fuel_type_id": "0d2209df-1000-4f29-9cc5-52aba3406fb3" ,
        "registration_manual": {
    "url": "https://bgm-s3.innovasive.dev/raw/parking/permit/registration_manual/20240220_db44ee1e-e2fa-41b9-b3bb-f000bd143f30_529783291.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=app%2F20240220%2Fap-southeast-1%2Fs3%2Faws4_request&X-Amz-Date=20240220T071325Z&X-Amz-Expires=10800&X-Amz-SignedHeaders=host&X-Amz-Signature=7a6c06f5ce3bac091831554185d754557de0d2a15c352409ed3135f755d20509",
    "content_type": "image/jpeg",
    "size": 250061,
    "original_filename": "side_eye_cat.jpg",
    "note": ""
}
        }
    ]
}
```







---

### [POST] /v1/app/parking/permits/:parking_id/car

**Method**: POST

**RequestURL**: `<<BGM_FACILITY>>/v1/parking/permit/:parking_id/car`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`page`</td>
    <td><code>int</code></td>
    <td>`1`</td>
    </tr>
    <tr>
    <td>`per_page`</td>
    <td><code>int</code></td>
    <td>`10`</td>
    </tr>
    <tr>
    <td>`statuses`</td>
    <td><code>string</code></td>
    <td>`VERIFIED`</td>
    </tr>
    <tr>
    <td>`statuses`</td>
    <td><code>string</code></td>
    <td>`REQUESTED`</td>
    </tr>
    <tr>
    <td>`is_expired`</td>
    <td><code>bool</code></td>
    <td>`false`</td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: `application/json`
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>parking_id</code></td>
    <td><code>a32070d1-7ac6-4cd4-8c5f-088f7bf62cd7</code></td>
    </tr>
</table>

**Request Body**:
```json
{
    "cars" : [
        {
        "id": "2a3067cf-0794-4ca2-aca4-a9b506ce02ce",
        "brand": "b1",
        "model": "m1",
        "color": "c1",
        "year": "2567",
        "capacity": "1500",
        "license_number": "กก2asf023",
        "province": "กรุงเทพมหานคร",
        "fuel_type_id": "0d2209df-1000-4f29-9cc5-52aba3406fb3" ,
        "registration_manual": {
    "url": "https://bgm-s3.innovasive.dev/raw/parking/permit/registration_manual/20240220_db44ee1e-e2fa-41b9-b3bb-f000bd143f30_529783291.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=app%2F20240220%2Fap-southeast-1%2Fs3%2Faws4_request&X-Amz-Date=20240220T071325Z&X-Amz-Expires=10800&X-Amz-SignedHeaders=host&X-Amz-Signature=7a6c06f5ce3bac091831554185d754557de0d2a15c352409ed3135f755d20509",
    "content_type": "image/jpeg",
    "size": 250061,
    "original_filename": "side_eye_cat.jpg",
    "note": ""
}
        }
    ]
}
```







---

### [PUT] /v1/parking/permit/:parking_id/status

**Method**: PUT

**RequestURL**: `<<BGM_FACILITY>>/v1/parking/permit/:parking_id/status`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: `application/json`
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>parking_id</code></td>
    <td><code>15d247b5-04f2-43c9-9807-c5a1ad46b646</code></td>
    </tr>
</table>

**Request Body**:
```json
{
    "status": "APPROVED",
    "parking_permit_type_year": "2567",
    "request_regis_amount": 5,
    "expire_at": "2025-02-07 23:59:00",
    "description": "สีชมพู",
    "gates": ["acd0bdc5-1367-4e48-a4cd-b05c219d64aa","625b555f-961c-47bd-b556-c3c0619a1dba","d4ef7306-aab9-4267-b479-3dc94f243849"],
    "buildings": ["fa96669c-ec8f-43c2-bde4-5555d1e4c4c7","9eca30a9-0514-4d6b-8046-515dfe463bf3"]
}
```







---

### [GET] /v1/gates

**Method**: GET

**RequestURL**: `<<BGM_FACILITY>>/v1/gates`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:






---

### [GET] /v1/parking/permit/status

**Method**: GET

**RequestURL**: `<<BGM_FACILITY>>/v1/parking/permit/status`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:






---

### [POST] /v1/parking/permits/list

**Method**: POST

**RequestURL**: `<<BGM_FACILITY>>/v1/parking/permits/list`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: `application/json`

**Request Body**:
```json
{
    "permits": [
        {
            "user_type_id": "88c8a7f6-2fe0-459c-a54d-0fd8564a1f0e",
            "organize_id": "fb69a17b-5cdd-4159-ba7e-4745d8c9c032",
            "contact_number": "088-888-1111",
            "contact_email": "example@kmutt.ac.th",
            "payment_proof": {
                "url": "https://bgm-s3.innovasive.dev/raw/20240113_86fc3d35-36b2-4569-94d1-b1ccfd99c3f9_745240936.jpg",
                "content_type": "image/jpeg",
                "size": 10150,
                "note": ""
            },
            "created_by": "03a3af4f-1881-4b6f-9b52-4933e8a8e73f"
        },{
            "user_type_id": "88c8a7f6-2fe0-459c-a54d-0fd8564a1f0e",
            "organize_id": "fb69a17b-5cdd-4159-ba7e-4745d8c9c032",
            "contact_number": "088-888-2222",
            "contact_email": "example@kmutt.ac.th",
            "payment_proof": {
                "url": "https://bgm-s3.innovasive.dev/raw/20240113_86fc3d35-36b2-4569-94d1-b1ccfd99c3f9_745240936.jpg",
                "content_type": "image/jpeg",
                "size": 10150,
                "note": ""
            },
            "created_by": "3d44203c-c4d4-4d1d-99bd-5d3315570bab"
        }
    ]
}
```







---

### [GET] /v1/receipt/:parking_id

**Method**: GET

**RequestURL**: `<<BGM_FACILITY>>/v1/receipt/:parking_id`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`secret`</td>
    <td><code>string</code></td>
    <td>`XXAKSJP2024`</td>
    </tr>
    </table>
**Authentication Type**: `none` 
    
**ContentType**: ``
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>parking_id</code></td>
    <td><code>e1e0c9ae-6e0e-4322-80ea-094073791653</code></td>
    </tr>
</table>

**Request Body**:






---

### [GET] /v1/receipt/:parking_id/file

**Method**: GET

**RequestURL**: `<<BGM_FACILITY>>/v1/receipt/:parking_id/file`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>parking_id</code></td>
    <td><code>e1e0c9ae-6e0e-4322-80ea-094073791653</code></td>
    </tr>
</table>

**Request Body**:






---

### [GET] /v1/parking/permits/list/excel

**Method**: GET

**RequestURL**: `<<BGM_FACILITY>>/v1/parking/permits/list/excel`


**Authentication Type**: `none` 
    
**ContentType**: ``

**Request Body**:






---

### [GET] /v1/parking/permits/stat/excel

**Method**: GET

**RequestURL**: `<<BGM_FACILITY>>/v1/parking/permits/stat/excel`


**Authentication Type**: `none` 
    
**ContentType**: ``

**Request Body**:






---

### [GET] /v1/parking/permits - Duplicate

**Method**: GET

**RequestURL**: `<<BGM_FACILITY>>/v1/parking/permits`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`page`</td>
    <td><code>int</code></td>
    <td>`1`</td>
    </tr>
    <tr>
    <td>`per_page`</td>
    <td><code>int</code></td>
    <td>`10`</td>
    </tr>
    <tr>
    <td>`statuses`</td>
    <td><code>string</code></td>
    <td>`VERIFIED`</td>
    </tr>
    <tr>
    <td>`statuses`</td>
    <td><code>string</code></td>
    <td>`REQUESTED`</td>
    </tr>
    <tr>
    <td>`is_expired`</td>
    <td><code>bool</code></td>
    <td>`false`</td>
    </tr>
    <tr>
    <td>`created_start_date`</td>
    <td><code>string</code></td>
    <td>`2024-02-06`</td>
    </tr>
    <tr>
    <td>`created_end_date`</td>
    <td><code>string</code></td>
    <td>`2024-02-06`</td>
    </tr>
    <tr>
    <td>`permit_type_codes`</td>
    <td><code>string</code></td>
    <td>`ORANGE`</td>
    </tr>
    <tr>
    <td>`permit_type_codes`</td>
    <td><code>string</code></td>
    <td>`PINK`</td>
    </tr>
    <tr>
    <td>`permit_type_codes`</td>
    <td><code>string</code></td>
    <td>`YELLOW`</td>
    </tr>
    <tr>
    <td>`statuses`</td>
    <td><code>string</code></td>
    <td>`REJECTED`</td>
    </tr>
    <tr>
    <td>`mode`</td>
    <td><code>string</code></td>
    <td>`STAFF`</td>
    </tr>
    <tr>
    <td>`create_status`</td>
    <td><code>string</code></td>
    <td>`RENEW`</td>
    </tr>
    <tr>
    <td>`user_type_id`</td>
    <td><code>uuid</code></td>
    <td>`92fe6591-7817-4c20-9c83-473798070cb5`</td>
    </tr>
    <tr>
    <td>`org_id`</td>
    <td><code>uuid</code></td>
    <td>`f53c0bee-c76f-4f15-939e-bbd9432bc36a`</td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:






---

## Folder: Internal
**Authentication**: `inherit`<p></p>



---

### Folder: SCHEDULER
**Authentication**: `inherit`<p></p>

---

#### [POST] sync all organize kmutt

**Method**: POST

**RequestURL**: `<<BGM_SCHEDULER>>/v1/scheduler/triggers`


**Authentication Type**: `none` 
    
**ContentType**: `application/json`

**Request Body**:

```json

{
    "name": "dag_sync_organize",
    "execute_datetime": "2023-05-31T20:19:00+07:00",
    "config": null
}

```



---

#### [POST] publish notification timer

**Method**: POST

**RequestURL**: `<<BGM_SCHEDULER>>/v1/scheduler/triggers`


**Authentication Type**: `none` 
    
**ContentType**: `application/json`

**Request Body**:

```json

{
    "name": "dag_sync_organize",
    "execute_datetime": "2023-05-31T20:19:00+07:00",
    "config": null
}

```



---

#### [GET] publish notification timer Copy

**Method**: GET

**RequestURL**: `<<BGM_SCHEDULER>>/v1/job/:job_id`


**Authentication Type**: `none` 
    
**ContentType**: `application/json`
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>job_id</code></td>
    <td><code></code></td>
    </tr>
</table>

**Request Body**:

```json

```








---

## Folder: NOTIFICATION
**Authentication**: `inherit`<p></p>**Headers**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Value</th>
    </tr><tr>
    <td>`<<API_KEY_NAME>>`</td>
    <td>`<<API_KEY_VALUE>>`</td>
    </tr></table>



---

### Folder: announcements
**Authentication**: `inherit`<p></p>

---

#### [GET] /v1/management/announcements

**Method**: GET

**RequestURL**: `<<BGM_NOTIFICATION>>/v1/management/announcements`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`search_word`</td>
    <td><code>unknow</code></td>
    <td> </td>
    </tr>
    <tr>
    <td>`publish_date`</td>
    <td><code>unknow</code></td>
    <td> </td>
    </tr>
    <tr>
    <td>`start_date`</td>
    <td><code>unknow</code></td>
    <td> </td>
    </tr>
    <tr>
    <td>`end_date`</td>
    <td><code>unknow</code></td>
    <td> </td>
    </tr>
    <tr>
    <td>`status`</td>
    <td><code>unknow</code></td>
    <td> </td>
    </tr>
    <tr>
    <td>`page`</td>
    <td><code>int</code></td>
    <td>`1`</td>
    </tr>
    <tr>
    <td>`per_page`</td>
    <td><code>int</code></td>
    <td>`20`</td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:



---

#### [GET] /v1/management/:announcement_id/announcements

**Method**: GET

**RequestURL**: `<<BGM_NOTIFICATION>>/v1/management/:announcement_id/announcements`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>announcement_id</code></td>
    <td><code>042420ed-3b03-4938-9b30-9c3607a303ad</code></td>
    </tr>
</table>

**Request Body**:



---

#### [POST] /v1/management/announcements

**Method**: POST

**RequestURL**: `<<BGM_NOTIFICATION>>/v1/management/announcements`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: `application/json`

**Request Body**:

```json

{
    "title": "test_scheduler",
    "publish_type": "CUSTOM", // NOW || CUSTOM
    "publish_date": "2024-04-10 15:30:00", // set null when publish_type == NOW
    "status": "DRAFT", // DRAFT || PUBLISH
    "detail": "test_scheduler",
    "cover": {
        "url": "https://bgm-s3.innovasive.dev/raw/cover/20240410_fbb82519-13e1-4d74-b116-475269d5d6aa_972460679.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=app%2F20240410%2Fap-southeast-1%2Fs3%2Faws4_request&X-Amz-Date=20240410T024835Z&X-Amz-Expires=10800&X-Amz-SignedHeaders=host&X-Amz-Signature=7d1338a126184e1164d596861934a28a9b3eceed3f792eaf4f755c5595c0722a",
        "content_type": "image/jpeg",
        "size": 55659,
        "original_filename": "gundam_aerial.jpg",
        "note": ""
    },
    "document": [
        {
        "url": "https://bgm-s3.innovasive.dev/raw/document/20240410_6f553912-cc39-4d9b-a79f-2a413f8550c5_622778745.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=app%2F20240410%2Fap-southeast-1%2Fs3%2Faws4_request&X-Amz-Date=20240410T024824Z&X-Amz-Expires=10800&X-Amz-SignedHeaders=host&X-Amz-Signature=642cd827856d24490efa91a4b638c45277495e8a6302e151f58126a9ee776180",
        "content_type": "image/jpeg",
        "size": 55659,
        "original_filename": "gundam_aerial.jpg",
        "note": ""
    }
    ]
}

```



---

#### [PUT] /v1/management/:announcement_id/announcements

**Method**: PUT

**RequestURL**: `<<BGM_NOTIFICATION>>/v1/management/:announcement_id/announcements`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: `application/json`
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>announcement_id</code></td>
    <td><code>042420ed-3b03-4938-9b30-9c3607a303ad</code></td>
    </tr>
</table>

**Request Body**:

```json

{
    "id": "042420ed-3b03-4938-9b30-9c3607a303ad",
    "title": "TEST COUNTER 2",
    "publish_type": "CUSTOM",
    "publish_date": "2024-04-10 18:01:00",
    "status": "WAIT",
    "detail": "<p>TEST COUNTER 2</p>",
    "total_view": 0,
    "updated_by": "ศุภกร ฉัตรอโณทัย",
    "updated_at": "2024-04-10 17:52:23",
    "cover": {
        "url": "https://bgm-s3.innovasive.dev/announcements/cover/20240410_4e2fa9c5-da25-4041-89a9-2afd802b4f4f_674716421.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=app%2F20240410%2Fap-southeast-1%2Fs3%2Faws4_request&X-Amz-Date=20240410T105916Z&X-Amz-Expires=10800&X-Amz-SignedHeaders=host&X-Amz-Signature=0546101500f0069a1d237a850a165f5cd8e00cde5690278ef03f44e6a737d4bc",
        "original_filename": "",
        "content_type": "image/jpeg",
        "size": 82208
    },
    "document": []
}

```



---

#### [DELETE] /v1/management/:announcement_id/announcements

**Method**: DELETE

**RequestURL**: `<<BGM_NOTIFICATION>>/v1/management/:announcement_id/announcements`


**Authentication Type**: `none` 
    
**ContentType**: `application/json`
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>announcement_id</code></td>
    <td><code></code></td>
    </tr>
</table>

**Request Body**:

```json

```



---

#### [GET] /v1/announcements

**Method**: GET

**RequestURL**: `<<BGM_NOTIFICATION>>/v1/announcements`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`search_word`</td>
    <td><code>unknow</code></td>
    <td> </td>
    </tr>
    <tr>
    <td>`page`</td>
    <td><code>int</code></td>
    <td>`1`</td>
    </tr>
    <tr>
    <td>`per_page`</td>
    <td><code>int</code></td>
    <td>`20`</td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:



---

#### [GET] /v1/:announcement_id/announcements

**Method**: GET

**RequestURL**: `<<BGM_NOTIFICATION>>/v1/:announcement_id/announcements`


**Authentication Type**: `none` 
    
**ContentType**: ``
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>announcement_id</code></td>
    <td><code></code></td>
    </tr>
</table>

**Request Body**:








---

### [GET] /v1/notification

**Method**: GET

**RequestURL**: `<<BGM_NOTIFICATION>>/v1/notification`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`page`</td>
    <td><code>int</code></td>
    <td>`1`</td>
    </tr>
    <tr>
    <td>`per_page`</td>
    <td><code>int</code></td>
    <td>`10`</td>
    </tr>
    <tr>
    <td>`title`</td>
    <td><code>string</code></td>
    <td>`Test`</td>
    </tr>
    <tr>
    <td>`publish_date`</td>
    <td><code>unknow</code></td>
    <td> </td>
    </tr>
    <tr>
    <td>`status`</td>
    <td><code>string</code></td>
    <td>`PUBLISH`</td>
    </tr>
    <tr>
    <td>`source`</td>
    <td><code>string</code></td>
    <td>`parking`</td>
    </tr>
    <tr>
    <td>`publish_start_date`</td>
    <td><code>unknow</code></td>
    <td> </td>
    </tr>
    <tr>
    <td>`publish_end_date`</td>
    <td><code>unknow</code></td>
    <td> </td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:






---

### [GET] /v1/notification/details

**Method**: GET

**RequestURL**: `<<BGM_NOTIFICATION>>/v1/notification/details`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`id`</td>
    <td><code>uuid</code></td>
    <td>`ba4669ea-a95a-4925-8ee9-81583c44d54b`</td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:






---

### [GET] /v1/notification/view

**Method**: GET

**RequestURL**: `<<BGM_NOTIFICATION>>/v1/notification`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`page`</td>
    <td><code>int</code></td>
    <td>`1`</td>
    </tr>
    <tr>
    <td>`per_page`</td>
    <td><code>int</code></td>
    <td>`10`</td>
    </tr>
    <tr>
    <td>`title`</td>
    <td><code>string</code></td>
    <td>`Test`</td>
    </tr>
    <tr>
    <td>`publish_date`</td>
    <td><code>unknow</code></td>
    <td> </td>
    </tr>
    <tr>
    <td>`status`</td>
    <td><code>string</code></td>
    <td>`PUBLISH`</td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:






---

### [GET] /v1/my/notification

**Method**: GET

**RequestURL**: `<<BGM_NOTIFICATION>>/v1/my/notification`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:






---

### [POST] /v1/notification

**Method**: POST

**RequestURL**: `<<BGM_NOTIFICATION>>/v1/notification`


**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: `application/json`

**Request Body**:
```json
{
    
    "title": "custom publish time",
    "status": "DRAFT",
    "publish_type": "NOW",
    "recipient_type": "ALL",
    "recipient": [],
    "detail": "<p>ทดสอบ PUBLISH CUSTOM + ALL</p>",
    "cover": {
        "url": "https://bgm-s3.innovasive.dev/raw/cover/20240329_0d0045cb-9606-409a-8869-8a8ba786cf19_225709681.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=app%2F20240329%2Fap-southeast-1%2Fs3%2Faws4_request&X-Amz-Date=20240329T074158Z&X-Amz-Expires=10800&X-Amz-SignedHeaders=host&X-Amz-Signature=8a3c6070bbeddfe800516793f492186d5fb0c85c50bbba5b2ddc9ec950306644",
        "content_type": "image/jpeg",
        "size": 40446,
        "original_filename": "Smiling-Cat.jpg",
        "note": ""
    },
    "publish_date": ""
}
```







---

### [DELETE] /v1/my/notification

**Method**: DELETE

**RequestURL**: `<<BGM_NOTIFICATION>>/v1/my/notification`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`notification_id`</td>
    <td><code>unknow</code></td>
    <td> </td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:






---

### [DELETE] /v1/notification

**Method**: DELETE

**RequestURL**: `<<BGM_NOTIFICATION>>/v1/notification`


**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>`notification_id`</td>
    <td><code>uuid</code></td>
    <td>`0b753e88-cf05-457d-a034-a6e4f149349d`</td>
    </tr>
    </table>
**Authentication Type**: `bearer` 
    

**BearerToken**: `<<JWT_TOKEN>>`

**ContentType**: ``

**Request Body**:






---

### [PUT] /v1/notification/:notification_id

**Method**: PUT

**RequestURL**: `<<BGM_NOTIFICATION>>/v1/notification/<<NOTIFICATION_ID>>`


**Headers**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Value</th>
    </tr>
    <tr>
    <td>Authorization</td>
    <td>`Bearer <<JWT_TOKEN>>`</td>
    </tr></table>
**Authentication Type**: `inherit` 
    
**ContentType**: `application/json`
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    <tr>
    <td><code>NOTIFICATION_ID</code></td>
    <td><code>0b753e88-cf05-457d-a034-a6e4f149349d</code></td>
    </tr>
    <tr>
    <td><code>RESPONSE</code></td>
    <td><code>{  "message": "no permission"}</code></td>
    </tr>
</table>

**Request Body**:
```json
{
    
    "title": "Only meeee",
    "status": "DRAFT",
    "publish_type": "NOW",
    "recipient_type": "UPLOAD",
    "recipient": ["c69a45d1-55af-41e2-9aa8-6e9b380fdd33"],
    "detail": "<p>ทดสอบ PUBLISH CUSTOM + ALL</p>",
    "cover": {
        "url": "https://bgm-s3.innovasive.dev/raw/cover/20240329_0d0045cb-9606-409a-8869-8a8ba786cf19_225709681.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=app%2F20240329%2Fap-southeast-1%2Fs3%2Faws4_request&X-Amz-Date=20240329T074158Z&X-Amz-Expires=10800&X-Amz-SignedHeaders=host&X-Amz-Signature=8a3c6070bbeddfe800516793f492186d5fb0c85c50bbba5b2ddc9ec950306644",
        "content_type": "image/jpeg",
        "size": 40446,
        "original_filename": "Smiling-Cat.jpg",
        "note": ""
    },
    "publish_date": ""
}
```







---



