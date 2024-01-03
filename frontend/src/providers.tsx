import { Datagrid, List, TextField, EditButton } from 'react-admin';

export const ProviderList = () => (
    <List>
        <Datagrid rowClick="show">
            <TextField source="id" />
            <TextField source="name" />
            <EditButton />
        </Datagrid>
    </List>
);