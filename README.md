# Go Contact
This is a simple contact application. This API is an exercise to learn Golang with PostgreSQL.

### Authentication
- User can register for a free account (Normal and Social e.g. Facebook)
- User can login to their account
- Authenticated user can create their contacts
- Authenticated user can get their contacts


#### Errors and Status Codes
If a request fails any validations, expect errors in the following format:

#### Testing API endpoints
<table>
<tr><th>Test</th>
<th>API-endpoint</th>
<th>HTTP-Verbs</th>
</tr>
<tr>
<td>Sign up a user</td>
<td>/api/user/new</td>
<td>POST</td>
</tr>
<tr>
<td>Login a user</td>
<td>/api/user/login</td>
<td>POST</td>
</tr>
<tr>
<td>Create a contact</td>
<td>/api/contacts/new</td>
<td>POST</td>
</tr>
<tr>
<td>Get contacts of a user</td>
<td>/api/user{id}/contacts</td>
<td>GET</td>
</tr>
</table>

### Authors
Arnold M. Osoro - mmosoroohh
