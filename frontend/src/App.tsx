import {
  Admin,
  Resource,
  ShowGuesser,
} from "react-admin";
import { dataProvider } from "./dataProvider";
import { TransactionList, TransactionEdit, TransactionCreate } from "./transactions";
import { AccountList } from "./accounts";
import { CostCenterList } from "./cost-centers"
import { ProviderList } from "./providers"
import PaidIcon from '@mui/icons-material/Paid';
import PersonIcon from '@mui/icons-material/Person';
import { authProvider } from './authProvider';

export const App = () => <Admin dataProvider={dataProvider}  authProvider={authProvider} requireAuth > 
    <Resource 
    name="transactions" 
    list={TransactionList} 
    edit={TransactionEdit} 
    create={TransactionCreate} 
    icon={PaidIcon} 
  />
  <Resource name="accounts" list={AccountList} show={ShowGuesser} recordRepresentation="description" />
  <Resource name="cost-centers" list={CostCenterList}  show={ShowGuesser}     recordRepresentation="description" />
  <Resource name="providers" list={ProviderList}  show={ShowGuesser}   recordRepresentation="name" icon={PersonIcon} />

  
</Admin>;
