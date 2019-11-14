# APC Class

## Get all Classes
* HTTP Request : ```GET http://api.com/class```
* Return a list of object in json format as follow

    ``` 
        [
            {
                "id"                    :   ObjectId,
                "professorfirstname"    :   String,
                "professorlastname"     :   String,
                "classname"             :   String,
                "address"               :   String,
                "year"                  :   Integer,
                "season"                :   Integer,
                "contestsids"           :   []Integer,
                "groupid"               :   String

            },...
        ]
    ```

## Get all Classes from a professor
* HTTP Request : ```GET http://api.com/class/{professorid}```
* Return a list of object in json format as follow

    ``` 
    [
        {
            "ID"                    :   String,
            "ProfessorID"           :   String,
            "professorfirstname"    :   String,
            "professorlastname"     :   String,
            "classname"             :   String,
            "address"               :   String,
            "year"                  :   Integer,
            "season"                :   Integer,
            "contestsids"           :   []Integer,
            "groupid"               :   String
        },...
    ]
    ```


## Create Classes
* HTTP Request : ```POST http://api.com/class```
* Send Classes's data in the request body in the following format
* *PS* : Only admin with `Professor` : `True` can make this request 

	``` 
    [
        {
            "professorfirstname"    :   String,
            "professorlastname"     :   String,
            "classname"             :   String,
            "address"               :   String,
            "year"                  :   Integer,
            "season"                :   Integer,
            `"contestsids"           :   []Integer,
            "groupid"               :   String`
        },...
    ]
    ``
* http StatusCreated (201) will be sent if the class has been created correctly


## Update Classes
* HTTP Request : ```PUT http://api.com/class```
* Send data in the request body in the following format
* *PS* : Only admin with `Professor` : `True` can make this request

    ``` 
    [
        {
            "id"                    :   ObjectId,
            "professorfirstname"    :   String,
            "professorlastname"     :   String,
            "classname"             :   String,
            "address"               :   String,
            "year"                  :   Integer,
            "season"                :   Integer,
            "contestsids"           :   []Integer,
            "groupid"               :   String
        }...
    ]
    ```

* http StatusCreated (201) will be sent if the student has been updated correctly


## Delete Classes
* HTTP Request : ```DELETE http://api.com/class```
* Send data in the request body in the following format
* *PS* : Only admin with `Professor` : `True` can make this request

	``` 
    [
        {  
            "id" : ObjectId
        },...
    ]
	```
* http StatusOK (200) will be sent if the students have been deleted correctly
