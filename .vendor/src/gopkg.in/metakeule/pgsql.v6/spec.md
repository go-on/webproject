# pgsql  

[HTML](.html)  
[Text (Markdown)](.md)  
[JSON](.json)  

project `pgsql`  
company `Know GmbH`  

persons  

- `Marc René Arns` (mra)  

requested by

- mra  

## Overview

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `planning` 
| **last update**                  | `2014-02-14` 
| **deadline**                     | `2014-01-05` 
| **est. working hours**           | `1` 
| **UUID**                         | `2170b38e-eb4f-4352-a7be-8229872a79a1` 


pgsql facilitates operations on postgresql databases in go

******
******

## Non-Goal


### N1 translations of error / validation messages

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `planning` 
| **last update**                  | `2014-01-05` 
| **UUID**                         | `74486739-2350-4597-899a-5f5a68849420` 


error messages / validation messages should not be translated via the
pgsql library.

instead the website should offer a way to get 

- all translation messages (for small projects)
- a translation for a given message
- translations for a given number of messages

and the clientside code (javascript) should query the translations 
from the website and show them instead of the original messages.

the translation service should also take a context to offer different 
translations in different contexts. the context is a URL optionally
followed by a hash. That hash may refer to an id of an element (e.g.
form element) or an anchor.

## Undecided


### U1 CRUD methoden und rest urls

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2014-01-27` 
| **UUID**                         | `92c37b03-985f-4da6-bd9d-04a439e642df` 


Es wird nur ein sinnvoller Teil der REST API Methoden 
(<http://en.wikipedia.org/wiki/Representational_state_transfer>) 
implementiert.

Das Format ist immer JSON und für eine Ressource unter 
`/api/v1/person` sind dann folgende Routen erzeugbar:


#### LIST

    GET /api/v1/person
    
Gibt eine Json Liste aller Personen

#### CREATE

    POST /api/v1/person
    
Erzeugt eine neue Person und gibt im Location header die URL
der neuen Ressource zurück. Im JSON steht der Primary Key  (ID) mit
dem entsprechenden neuen Wert

#### READ

    GET /api/v1/person/1
    
Gibt die JSON Respresentation der Person mit der ID 1 zurück

#### UPDATE

    PUT /api/v1/person/1
    
Aktualisiert die Person mit der ID 1 entsprechend den übergebenen 
Werten

#### delete

    delete /api/v1/person/1

******
******

### U2 Abgleich mit heroku API

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `` 
| **state**                        | `` 
| **last update**                  | `2014-01-18` 
| **UUID**                         | `3daef0cb-e3cf-4065-91c8-5c0f7d67825c` 


URL: <https://devcenter.heroku.com/articles/platform-api-reference#overview>  
Die API von Heroku ist ein guter Orientierungspunkt für Fragen, wie

- Bereitstellung div. API Versionen
- `CRUD / REST`  
          

     `delete` used for destroying existing objects  
     `GET` 	used for retrieving lists and individual objects  
     `HEAD` 	used for retrieving metadata about existing objects  
     `PATCH` 	used for updating existing objects  
     `PUT` 	used for replacing existing objects  
     `POST` 	used for creating new objects  

- Authentifizierung
- Caching (das `ETag` könnte eine Prüfsumme über das zurückgegebene 
  Ergebnis sein und wird dann mit `If-None-Match` header abgeglichen)
- `JSON schema` <https://blog.heroku.com/archives/2014/1/8/json_schema_for_heroku_platform_api>
- custom types  

    `date-time` 	string 	timestamp in iso8601 format  
    `uuid` 	string 	uuid in 8-4-4-4-12 format  
    
- Error Responses
- Method Override  
  When using a client that does not support all of the methods, 
  you can override by using a `POST` and setting the 
  `X-Http-Method-Override` header to the desired methed.  
  For instance, to do a `PATCH` request, do a `POST` with header 
  `X-Http-Method-Override: PATCH`.
- Data integrity  
  You may pass the If-Match header with an ETag value from a previous 
  response to ensure a resource has not changed since you last 
  received it. If the resource has changed, you will receive a 412 
  Precondition Failed response. If the resource has not changed, the 
  request will proceed normally.
- Parameters  
  Values that can be provided for an action are divided between 
  optional and required values. The expected type for each value is
  specified and unlisted values should be considered immutable. 
  Parameters should be JSON encoded and passed in the request body.
- Ranges 
  List requests will return a Content-Range header indicating the 
  range of values returned. Large lists may require additional 
  requests to retrieve. If a list response has been truncated you 
  will receive a 206 Partial Content status and one or both of 
  Next-Range and Prev-Range headers if there are next and previous 
  ranges respectively. To retrieve the next or previous range, 
  repeat the request with the Range header set to either the 
  Next-Range or Prev-Range value from the previous request.  
  The number of values returned in a range can be controlled using 
  a max key in the Range header. For example, to get only the first 
  10 values, set this header: Range: id ..; max=10;. max can also be 
  passed when iterating over Next-Range and Prev-Range. The default 
  page size is 200 and maximum page size is 1000.  
  The property used to sort values in a list response can be changed. 
  The default property is id, as in Range: id ..;. To learn what other
  properties you can use to sort a list response, inspect the 
  Accept-Ranges headers. For the apps resource, for example, you can 
  sort on either id or name: `Accept-Ranges: id, name`
- The default sort order for resource lists responses is ascending. 
  You can opt for descending sort order by passing a order key in the 
  range header: `Range: name ..; order=desc;`  
  Combining with the max key would look like this:  
  `Range: name ..; order=desc,max=10;`
  
##### Successful Responses

    Status 	Description
    
    200 ok 	request succeeded
    
    201 Created 	resource created, for example a new app was created or an add-on was provisioned
    
    202 Accepted 	request accepted, but the processing has not been completed
    
    206 Partial Content 	request succeeded, but this is only a partial response, see ranges


##### Error Responses

Error responses can be divided in to two classes. Client errors result from malformed requests and should be addressed by the client. Heroku errors result from problems on the server side and must be addressed internally.

###### Client Error Responses

    Status 	Error ID 	Description
    400 Bad Request 	bad_request 	request invalid, validate usage and try again
    401 Unauthorized 	unauthorized 	request not authenticated, validate credentials and try again
    402 Payment Required 	delinquent 	either the account has become delinquent as a result of non-payment, or the account’s payment method must be confirmed to continue
    403 Forbidden 	forbidden 	request not authorized, provided credentials do not provide access to specified resource
    403 Forbidden 	suspended 	request not authorized, account or application was suspended.
    404 Not Found 	not_found 	request failed, the specified resource does not exist
    406 Not Acceptable 	not_acceptable 	request failed, set Accept: application/vnd.heroku+json; version=3 header and try again
    416 Requested Range Not Satisfiable 	requested_range_not_satisfiable 	request failed, validate Content-Range header and try again
    422 Unprocessable Entity 	invalid_params 	request failed, validate parameters try again
    422 Unprocessable Entity 	verification_needed 	request failed, enter billing information in the Heroku Dashboard before utilizing resources.
    429 Too Many Requests 	rate_limit 	request failed, wait for rate limits to reset and try again, see rate limits

###### Heroku Error Responses

    Status 	Description
    500 Internal Server Error 	error occurred, we are notified, but contact support if the issue persists
    503 Service Unavailable 	API is unavailable, check response body or Heroku status for details


###### Oauth

<https://devcenter.heroku.com/articles/oauth>


###### Autogenerating go client for json schema based api:

<https://blog.heroku.com/archives/2014/1/9/auto_generating_a_go_api_client_for_heroku>

and

<https://github.com/bgentry/heroku-go>




******
******

### U3 neuimplementierung mit trennung zwischen sql erzeugung und responsebearbeitung

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `planning` 
| **last update**                  | `2014-01-29` 
| **UUID**                         | `9a7f567a-5416-4ce0-9728-5ac79aa0bacf` 


Die Idee ist, ein package `sql` zu haben, welches nur das Datenbank
spezifische SQL erzeugt und aber allgemeinen Interfaces genügt:


    github.com/go-on/sql/
       postgres/
       mysql/
       sqlite/
       
Ein solches Interface könnte z.B. sein:

    type SQL interface {
       SQL() string
    }
    
    type Field interface {
       SQL
       Table() SQL
       Name() string
       TableName() string
    }
    
Dann gäbe es für `Select`, `Update` usw. Funktionen, z.B.

    type pgSelect struct {
       Fields []Field
       Aliases []Alias
       Where []Condition
       ....
    }
    
    func (s *pgSelect) SQL() string {
       // compose the sql string
    }
    
    func SELECT(fields ...Field) pgSelect {
       
    }
    
Diese Funktionen hießen immer gleich. 
Funktionen gebe es auch für:

- INSERT
- UPDATE
- SELECT
- delete
- CREATE\_TABLE
- ALTER\_TABLE
- DROP\_TABLE

Dann ggf. methoden, z.B.

- UNION
- WHERE
- AND eine where bedingung die als Parameter einer anderen Bedingung
  übergeben wird und damit zu derem "AND" wird
- OR eine where bedingung, die als Parameter einer anderen Bedingung
  übergeben wird und damit zu derem "OR" wird
- GROUP\_BY
- ORDER\_BY
- JOIN
- LEFTJOIN
- RIGHTJOIN
- INNERJOIN


Eine selectabfrage könnte dann wie folgt aussehen:

    person := TableStr("person")
    firstname := FieldStr("person.firstname")
    lastname := FieldStr("person.lastname")
    company := FieldStr("person.company")
    companyid := FieldStr("company.id")
    companyname := FieldStr("company.name")

    // order of parameters is irrelevant
    sql := SELECT(firstname, lastname, companyname).
      FROM(person).WHERE(
        firstname, EQUALS, "Peter", 
        OR( lastname, MATCHES, "Me(y|i)er", 
          AND(firstname, EQUALS, "Susi"),
        ),
      ).
      ORDER_BY(lastname, ASC, firstname, DESC).
      LIMIT(12).
      OFFSET(2).
      JOIN(
        companyname,
        ON(company, companyid),
      ).SQL()

Die untersützten Typen samt typenumwandlung und das rückgeschreibe
der Werte wären Gegenstand anderer Bibliotheken.

## Definition


### D1 REST von Wikipedia

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2014-01-18` 
| **UUID**                         | `471b5863-5d23-4134-b420-f0c1d0531bd3` 


gemäß <http://de.wikipedia.org/wiki/Representational_State_Transfer>


#### GET

fordert die angegebene Ressource vom Server an. GET weist keine Nebeneffekte auf. Der Zustand am Server wird nicht verändert, weshalb GET als sicher bezeichnet wird.

#### POST

fügt eine neue (Sub-)Ressource unterhalb der angegebenen Ressource ein. Da die neue Ressource noch keine URI besitzt, adressiert den URI die übergeordnete Ressource. Als Ergebnis wird der neue Ressourcenlink dem Client zurückgegeben.

#### PUT

die angegebene Ressource wird angelegt. Wenn die Ressource bereits existiert, wird sie geändert.

#### PATCH

ein Teil der angegeben Ressource wird geändert. Hierbei sind Nebeneffekte erlaubt.

#### delete

löscht die angegebene Ressource. Wenn der Client versucht, eine Ressource zu löschen, die nicht existiert bzw. bereits gelöscht wurde, erhält der Client – sofern die REST-Schnittstelle korrekt implementiert wurde – keine Fehlermeldung (siehe auch: HTTP-Statuscodes). Abhängig von der Implementierung wird eine Ressource meist – entgegen der HTTP-Spezifikation – nicht physisch gelöscht, sondern nur als gelöscht gekennzeichnet und somit versteckt und deaktiviert.

#### HEAD

fordert Metadaten zu einer Ressource vom Server an.

#### OPTIONS

prüft, welche Methoden auf einer Ressource zur Verfügung stehen.


## Feature


### F1 Support for `If-Match` header wie heroku

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2014-02-14` 
| **UUID**                         | `b03408a7-8603-41ba-b328-3217fb30adb5` 


##### Data integrity

You may pass the If-Match header with an ETag value from a previous 
response to ensure a resource has not changed since you last 
received it. If the resource has changed, you will receive a 412 
Precondition Failed response. If the resource has not changed, the 
request will proceed normally.


sollte vielleicht auch als rack wrapper gegeben werden

******
******

### F2 Get und List sollten nur Keys zurückgeben, die auch in der Definition stehen

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2014-01-18` 
| **UUID**                         | `daac10b8-3dff-46cf-bb80-86ff0552fb7c` 


Get und List sollten nur Keys zurückgeben, die auch in der Definition 
stehen, dazu ist es nötig, vor dem Json export in ein map zu transformieren

******
******

### F3 List sollte ein leeres Json Array zurückgeben bei einer leeren Liste

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2014-01-18` 
| **UUID**                         | `b37b43ec-0153-41a5-b760-ceeffb8a100d` 




******
******

### F4 Get sollte 404 zurückgeben, wenn Id nicht vorhanden ist

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2014-01-18` 
| **UUID**                         | `8f8c70a5-170a-4df0-b2dc-2f4f9a0ad12c` 




******
******

### F5 HEAD unterstützen

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2014-02-08` 
| **UUID**                         | `eeb2b4fa-107d-4d6c-affb-e7f6272f971b` 


`HEAD` soll das gleiche zurückliefern wie `GET`, nur ohne Body

******
******

### F6 ETag unterstützen

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2014-02-08` 
| **UUID**                         | `b4b7aed0-6521-4182-8e83-3aa6b5941fcb` 


jede `GET` anfrage auf eine einzelne Ressource soll ein Etag als Prüfsumme
bilden und zurückgeben. (recherchieren, ob eine Prüfsumme genügt).


******
******

### F7 statt PUT, PATCH registrieren

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2014-02-08` 
| **UUID**                         | `854c2be1-68e2-4372-bd2f-3a0d33ac28ee` 


`PUT` wird für unseren Anwendungsfall nicht benötigt. Stattdessen `PATCH`.
Da man ja im Struct definieren kann, was alles über `PATCH` aktualisiert
werden kann, kann dort auch alles stehen.



******
******

### F8 Prüfen auf `If-None-Match` Header

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2014-02-14` 
| **UUID**                         | `e2ec7026-a4c8-48a3-84be-55bad3850121` 


Falls der Header gesetzt wird und das entsprechendes `ETag` gleich ist,
soll ein `PATCH` nicht ausgeführt werden.

sollte auch als rack wrapper gegeben werden

******
******

### F9 `X-Http-Method-Override` Header unterstützen

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2014-02-14` 
| **UUID**                         | `e1fcb4cc-1454-4e8a-a76f-b7d35a3b5c0e` 


ist über `rack/wrapper` gewährleistet
functioniert zur zeit mit router nicht (prüfen)

******
******

### F10 Success Types wie heroku

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2014-02-14` 
| **UUID**                         | `0647c48f-7ea7-4249-9836-e146202927e2` 


##### Successful Responses

    Status 	Description
    
    200 ok 	request succeeded
    
    201 Created 	resource created, for example a new app was created or an add-on was provisioned
    
    202 Accepted 	request accepted, but the processing has not been completed
    
    206 Partial Content 	request succeeded, but this is only a partial response, see ranges


******
******

### F11 Error States

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2014-02-14` 
| **UUID**                         | `d56c8077-568f-44f7-9315-2a8feab5c317` 


###### Client Error Responses

    Status 	Error ID 	Description
    400 Bad Request 	bad_request 	request invalid, validate usage and try again
    401 Unauthorized 	unauthorized 	request not authenticated, validate credentials and try again
    402 Payment Required 	delinquent 	either the account has become delinquent as a result of non-payment, or the account’s payment method must be confirmed to continue
    403 Forbidden 	forbidden 	request not authorized, provided credentials do not provide access to specified resource
    403 Forbidden 	suspended 	request not authorized, account or application was suspended.
    404 Not Found 	not_found 	request failed, the specified resource does not exist
    406 Not Acceptable 	not_acceptable 	request failed, set Accept: application/vnd.heroku+json; version=3 header and try again
    416 Requested Range Not Satisfiable 	requested_range_not_satisfiable 	request failed, validate Content-Range header and try again
    422 Unprocessable Entity 	invalid_params 	request failed, validate parameters try again
    422 Unprocessable Entity 	verification_needed 	request failed, enter billing information in the Heroku Dashboard before utilizing resources.
    429 Too Many Requests 	rate_limit 	request failed, wait for rate limits to reset and try again, see rate limits

###### Heroku Error Responses

    Status 	Description
    500 Internal Server Error 	error occurred, we are notified, but contact support if the issue persists
    503 Service Unavailable 	API is unavailable, check response body or Heroku status for details



******
******

### F12 Unterstützung für `OPTIONS` Anfrage

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2014-02-14` 
| **UUID**                         | `b05b9d1f-245a-4570-b1b8-e567d6db11b2` 


entsprechende `OPTIONS` anfrage mit 
maximal `Allow: GET HEAD PATCH delete`  auf ressources und maximal
`Allow: GET POST` auf listen urls. (Prüfen, ob die `Allow`-Syntax stimmt.

******
******

### F13 Missing CRUD Tests

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `agreed` 
| **last update**                  | `2014-01-27` 
| **UUID**                         | `a044ed8d-81bf-4009-935d-21dbe3060e35` 


Tests are missing for

#### Handler Creator Errors

- table has a composed primary key
- table has no primary key
- table is not registered
- table is not created
- delete tag on non primary key
- create create / update / list / read / delete handlers without
  having the a correspondig tag set
- field is not a (proper) pgsql field
- proper error types
  
#### Types

- handle all types supported by fat structs, except maps and slices
- raise error for map and slice types
- proper validation errors

#### Actions: Error Conditions

- wrong type of id
- wrong type of patch / post parameter
- missing post / patch parameter
- double post / patch parameter
- proper error types

#### Handlers

- Error Responses
- Success Responses for PATCH, POST and delete
- Status Codes
- response Headers for POST
- response content-types
- validations, single field validations for PATCH and POST
- proper distinguish validation errors and server errors
- allow error handler to be passed it and be used for server errors

******
******

### F14 CRUD call

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Implementing` 
| **last update**                  | `2014-01-27` 
| **UUID**                         | `dc490d15-117d-4b7d-86b1-facd92bbebc1` 


We want to able to define simple rest calls via fat structs, e.g.


    type Person struct {
	    Id        *fat.Field `[...] pgsql.rest:"delete,get,index"`
	    LastName  *fat.Field `[...] pgsql.rest:"get,post,put,index"`
	    Age       *fat.Field `[...] pgsql.rest:"get"`	
    }
    
    func (p *Person) Validate() (invalidFields []string, err error) {
    	// Do some validation...
    }
    
    var PERSON = fat.Proto(&Person{}).(*Person)
    
    func init() {
        r := router.New()
        pgsql.NewCRUD(PERSON).MountAll(db, r, "/person")
        http.ListenAndServe(":8080", r)
    }
    
    
This should create the following routes

    GET /person          // list of persons as json
    GET /person/:id      // single person as json
    PUT /person          // success in json
    delete /person/:id   // success in json
    POST /person/:id     // success in json

with their handlers. Each route should only respect the fields
with the corresponding `pgsql.crud` tag.

The following register methods should be available

    MountIndex()
    MountCREATE()
    MountREAD()
    MountUPDATE()
    Mountdelete()
    MountLIST()
    MountAll()
    Mount(CREATE|READ|UPDATE|LIST|delete)

That should allow to mount different rest handlers on different
routes.

**Protection of routes will be done via the middleware in the router.**

All validation and type and default value features of the fat structs
should apply.

`POST` and `PUT` should also respect a special header `X-Validation`
which results in no database action but only realistic validation
checking and error reporting. This header could be used in ajax forms
to validate and report validation errors before submitting a form.
If the addition header `X-Validate-Field` is set, only the given field
is validated. So the normal procedure for a javascript validation 
library would be:

- on blur of each field, submit the form with `X-Validation` set to
  true and `X-Validate-Field` set to the blurred field
  
- if their is a validation error, show the error near the field

- if the last field is blurred, submit the form with `X-Validation` set to
  true without `X-Validate-Field` to get the full validation
  
- if their is are validation errors, show the errors near the corresponding
  fields
  
- continue if the values of fields change to validate with full 
  validation  
  
- if there are no validation errors and missing fields, allow to submit
  the form, report any errors returned from submitting the form

******
******

### F15 CRUD list

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2014-02-14` 
| **UUID**                         | `d5ce0061-c406-4080-9bc1-2a4833f58fbc` 


voraussetzung ist, dass offset und limit übergeben werden können

Spezielle variante, die per json eine Liste abfragt.

Grundsätzlich wird ein `int` als `LIMIT` angegeben,
wenn das int `< 0` ist, werden alle zurückgegeben.

Als URL-Parameter werden berücksichtigt:

- limit (kann noch weniger einfordern als vordefiniertes Limit)
- offset
- sortby (liste von feldern, nach denen sortiert wird)
- sortorder (liste von sortierungsrichtungen, muss genausoviele 
  einträge haben, wie sortby)
  
  
#### so macht heroku es und sollten wir es auch machen:



#### Ranges 
  
  List requests will return a Content-Range header indicating the 
  range of values returned. Large lists may require additional 
  requests to retrieve. If a list response has been truncated you 
  will receive a 206 Partial Content status and one or both of 
  Next-Range and Prev-Range headers if there are next and previous 
  ranges respectively. To retrieve the next or previous range, 
  repeat the request with the Range header set to either the 
  Next-Range or Prev-Range value from the previous request.  
  The number of values returned in a range can be controlled using 
  a max key in the Range header. For example, to get only the first 
  10 values, set this header: Range: id ..; max=10;. max can also be 
  passed when iterating over Next-Range and Prev-Range. The default 
  page size is 200 and maximum page size is 1000.  
  The property used to sort values in a list response can be changed. 
  The default property is id, as in Range: id ..;. To learn what other
  properties you can use to sort a list response, inspect the 
  Accept-Ranges headers. For the apps resource, for example, you can 
  sort on either id or name: `Accept-Ranges: id, name`
- The default sort order for resource lists responses is ascending. 
  You can opt for descending sort order by passing a order key in the 
  range header: `Range: name ..; order=desc;`  
  Combining with the max key would look like this:  
  `Range: name ..; order=desc,max=10;`
  
 siehe auch: <http://benramsey.com/blog/2008/05/206-partial-content-and-range-requests/>
 
 siehe auch: <http://tools.ietf.org/html/rfc2616#section-10.4.17>

******
******

### F16 CRUD success responses

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2014-02-14` 
| **UUID**                         | `e12ed0ed-06eb-4c0c-9fc9-7b0b3cbe54bd` 


- **`LIST / GET`**: `Status 200`, Array von jsonifizierten Werten
- **`READ / GET`**: `Status 200` jsonifizierter Wert
- **`CREATE / POST`**: `Status 201`
  das `Location` Feld im Header gibt den Ort der neuen Ressource an
  der Body die Id der neuen Ressource 
- **`UPDATE / PUT`**: `Status 204`
- **`delete`**: `Status 204`
  keine Daten



******
******

### F17 CRUD fehlermeldungen

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Implementing` 
| **last update**                  | `2014-02-14` 
| **UUID**                         | `37cc8ec1-e2f6-4cc5-b7db-ffdbf654fa19` 


Grundsätzlich haben Fehlermeldungen vom Status `500` die Fehlermeldung
im Body.
    
Validierungsfehler sehen so aus (Status `422`):

    {
      "ValidationErrors": {
        "fieldname": "validation error message"
      },
      "Error": "invalid data"
    }
    
Die Validierungsfehler kommen nur bei `PUT` und `POST` Routen.

Wird die Resource nicht gefunden, so gibt es eine Meldung folgender
Art mit dem Status `404` und dem Body `Not found`

******
******

### F18 Validation

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `agreed` 
| **last update**                  | `2014-02-14` 
| **UUID**                         | `b8109600-d8ac-4137-afef-bf5015d187fa` 


`POST` and `PUT` should also respect a special header `X-Validation`
which results in no database action but only realistic validation
checking and error reporting. This header could be used in ajax forms
to validate and report validation errors before submitting a form.
If the addition header `X-Validate-Field` is set, only the given field
is validated. So the normal procedure for a javascript validation 
library would be:

- on blur of each field, submit the form with `X-Validation` set to
  true and `X-Validate-Field` set to the blurred field
  
- if their is a validation error, show the error near the field

- if the last field is blurred, submit the form with `X-Validation` set to
  true without `X-Validate-Field` to get the full validation
  
- if their is are validation errors, show the errors near the corresponding
  fields
  
- continue if the values of fields change to validate with full 
  validation  
  
- if there are no validation errors and missing fields, allow to submit
  the form, report any errors returned from submitting the form

******
******
******
******


