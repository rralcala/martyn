import { Datagrid, List, TextField, EditButton } from 'react-admin';

export const CostCenterList = () => (
    <List>
        <Datagrid rowClick="show">
            <TextField source="id" />
            <TextField source="description" />
            <EditButton />
        </Datagrid>
    </List>
);