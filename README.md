# packing-project

This project will accept a number of items and return the optimal packaging, it follows these constraints:
1. Only whole packs can be sent. Packs cannot be broken open.
2. Within the constraints of Rule 1 above, send out no more items than necessary to fulfil the order.
3.	Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfil each order.

https://packing-project.onrender.com

# Endpoints:

GET https://packing-project.onrender.com/pack/250
- Description: Retrieves packing details for a specified number.
- Path Parameter: number (integer) - The number for which packing details are requested.
- Response: JSON structure containing packing details such as ExtraItems, TotalItems, TotalPacks, and information about different Packs.
```JSON
    {
    "ExtraItems": 0,
    "TotalItems": 250,
    "TotalPacks": 1,
    "Packs": {
        "box_1": {
            "size": 5000,
            "used": 0
        },
        "box_2": {
            "size": 2000,
            "used": 0
        },
        "box_3": {
            "size": 1000,
            "used": 0
        },
        "box_4": {
            "size": 500,
            "used": 0
        },
        "box_5": {
            "size": 250,
            "used": 1
            }
        }
    }
```
- Usage: GET /pack/250 would retrieve packing details for the number 250.

---

POST https://packing-project.onrender.com/pack/add?size=15
- Description: Adds a new pack size to the system.
- Request Body: JSON structure containing the details of the pack size to be added (e.g., size of the pack).
- Response: Status 200 OK
- Usage: POST /pack/add with a JSON body specifying the new pack size details.

---

DELETE https://packing-project.onrender.com/pack/delete?size=15
- Description: Removes an existing pack size from the system.
- Request Body: JSON structure specifying the pack size to be deleted.
- Response: Status 200 OK
- Usage: DELETE /pack/delete with a JSON body specifying the pack size to remove.

---

GET https://packing-project.onrender.com/pack/list
- Description: Lists all available pack sizes in the system.
- Response: JSON array of all pack sizes, each with its detailed information.
```JSON
    {
        "pack_sizes": [
            250,
            500,
            1000,
            2000,
            5000
        ]
    }
```
- Usage: GET /pack/list to retrieve a list of all pack sizes.

---

POST https://packing-project.onrender.com/pack/new
- Description: Creates a new list of pack sizes, replacing the existing list.
- Request Body: JSON array containing the details of new pack sizes to be created.
```JSON
    {
        "pack_sizes": [
            200, 
            100, 
            500
        ]
    }
```
- Response: Confirmation message indicating successful creation of the new list of pack sizes.
- Usage: POST /pack/new with a JSON body specifying a new list of pack sizes.

---

GET https://packing-project.onrender.com/pack/reset
- Description: Resets the list of pack sizes to default settings or removes all custom pack sizes.
- Response: Status 200 OK
- Usage: GET /pack/reset to reset pack sizes to their default settings.

---
# Postman Collection
in /postman there is a postman collection with the endpoints and sample data ready to go!