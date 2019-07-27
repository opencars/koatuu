### /territories/{code}
---
##### ***GET***
**Summary:** Get information about territory by unique code.

**Responses**

| Code | Description  | Schema |
| ---- | ------------ | ------ |
| 200  | Territory    | [Territory](#territory)  |
| 404  | Error        | [Error](#error)          |

### /territories?level=1
---
##### ***GET***
**Summary:** Get list of level = 1 territories.

**Responses**

| Code | Description  | Schema |
| ---- | ------------ | ------ |
| 200  | Territory    | [ [Territory](#territory) ]  |
| 404  | Error        | [Error](#error)              |

### /territorie?level=2
---
##### ***GET***
**Summary:** Get list of level = 2 territories.

**Responses**

| Code | Description  | Schema |
| ---- | ------------ | ------ |
| 200  | Territory    | [ [Territory](#territory) ]  |
| 404  | Error        | [Error](#error)              |

### /territories?level=3
---
##### ***GET***
**Summary:** Get list of level = 3 territories.

**Responses**

| Code | Description  | Schema |
| ---- | ------------ | ------ |
| 200  | Territory    | [ [Territory](#territory) ]  |
| 404  | Error        | [Error](#error)              |

### Models
---

### Territory

| Name   | Type   | Description    |
| ------ | ------ | -------------- |
| code   | string | Unique code.   |
| name   | string | Description of the territory |
| type   | string | Type of the territory        |
| сhilds | [ Territory ] | List of child territories |

### Error

| Name  | Type   | Description    |
| ----- | ------ | -------------- |
| error | string | Error message. |
