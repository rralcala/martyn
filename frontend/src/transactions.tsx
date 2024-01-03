import {
    Create,
    Datagrid,
    DateField,
    DateInput,
    Edit,
    EditButton,
    List,
    NumberField,
    NumberInput,
    ReferenceField,
    ReferenceInput,
    required,
    SimpleForm,
    TextField,
    TextInput,
} from 'react-admin';

import { exporter }from './transactionExporter'

// in src/posts.tsx
const transactionFilters = [
   
    <ReferenceInput source="provider_id" label="Provider" reference="providers" />,
    <ReferenceInput source="cost_center_id" label="Cost Center" reference="cost-centers" />,
    <ReferenceInput source="account_id" label="Account" reference="accounts" />,
];

export const TransactionList = () => (
    <List filters={transactionFilters} exporter={exporter}>
        <Datagrid>
            <DateField source="date" />
            <ReferenceField source="provider" reference="providers"  link="show" />
            <TextField source="description" />
            <NumberField source="amount"/>
            <ReferenceField source="cost_center" reference="cost-centers"  link="show"/>
            <ReferenceField source="account" reference="accounts"  link="show" />
            <EditButton />
        </Datagrid>
    </List>
);

export const TransactionEdit = () => (
    <Edit>
        <SimpleForm>
            <TextInput disabled source="id" />
            <DateInput source="date" validate={[required()]}/>
            <ReferenceInput  source="provider" reference="providers" />
            <TextInput source="description"  multiline rows={3} validate={[required()]}/>
            <ReferenceInput  source="cost_center" reference="cost-centers" />
            <ReferenceInput  source="account" reference="accounts" />
        </SimpleForm>
    </Edit>
);

export const TransactionCreate = () => (
    <Create>
        <SimpleForm>
            <DateInput source="date"  validate={[required()]}/>
            <ReferenceInput  source="provider" reference="providers"  />
            <TextInput source="description"  multiline rows={3}  validate={[required()]}/>
            <NumberInput source="amount"  validate={[required()]} />
            <ReferenceInput  source="cost_center" reference="cost-centers" />
            <ReferenceInput  source="account" reference="accounts" />
        </SimpleForm>
    </Create>
);