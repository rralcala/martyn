
import jsonExport from 'jsonexport/dist';
import { downloadCSV }from 'react-admin';

export const exporter = async (transactions, fetchRelatedRecords) => {

    const accounts = await fetchRelatedRecords(transactions, 'account', 'accounts')
    const providers = await fetchRelatedRecords(transactions, 'provider', 'providers')
    const costCenters = await fetchRelatedRecords(transactions, 'cost_center', 'cost-centers')
   
    const postsForExport = transactions.map(transaction => {
        transaction.date = transaction.date.split('T')[0];
        const { id, provider, cost_center, account, ...postForExport } = transaction; // omit backlinks and author
        postForExport.provider = providers[transaction.provider].name; // add a field
        postForExport.account = accounts[transaction.account].description;
        postForExport.cost_center = costCenters[transaction.cost_center].description;
        return postForExport;
    });
    jsonExport(postsForExport, {
        headers: ['date', 'description', 'amount', 'provider', 'cost_center'] // order fields in the export
    }, (err, csv) => {
        downloadCSV(csv, 'movimientos'); // download as 'posts.csv` file
    });
};