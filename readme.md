```mermaid
flowchart LR
file2["Users
csv
100gb"
]

file1["Products
csv
200gb"
]

code["application"]

storage["Any storage"]

file1 --> code
file2 --> code

code --> storage

```