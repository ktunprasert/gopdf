# Gotenberg Invoice PDF Generator

1. Frontend for modifying logos and saving configs
2. Be able to generate 4 templates which have the same template but text slightly altered
3. Able to preview through API routes
4. Able to parse input data from Excel and automatically split into multiple sets of 4 templates when an item exceeds the table limit
    - i.e. If one PDF can manage upto 10 items then an excel sheet of 20 items must generate 2 sets of PDFs (one set is 4 documents mirrored for tax invoice and copies)

## Todo

1. Add in SQL module for being able to save multiple invoices and Logo paths
2. Supprt multiple tenants with their own configurationn
3. Uploading system for the clients to save their logos
4. Uploading Excel or spreadsheet will spawn multiple preview per input

## Access patterns

Single table design with partition key separation. Access each object by using the entity prefix and unique identifier. Design is outlined as follows

Tenants: tenant:{tenant\_name}
Invoice: {tenant\_name}:{invoice\_number}

For example, we have a two tenants called Beonit and Avocado using different types of invoice number and the data is structured as follows

| Partition | Unique ID |
| --------- | --------- |
| tenant     | tenant:beonit |
| tenant     | tenant:avocado |
| invoice    | invoice#beonit:1 |
| invoice    | invoice#beonit:2 |
| invoice    | invoice#avocado:10001 |
| invoice    | invoice#avocado:10002 |

We can retrieve the collection by calling the prefix url route on the datastore table and single item read/write by inferring from that collection. It should be said that the tenant name should be unique and will always be lowercase. If you know the tenant company entity already then you will be able to access the item directly calling the path using the composite key.
