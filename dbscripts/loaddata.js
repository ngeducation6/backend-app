const { MongoClient } = require('mongodb');

async function loadDatabase() {
//   const uri = 'mongodb://localhost:27017';
  const uri = 'mongodb://testdev:password123@localhost:27017/admin';
  const client = new MongoClient(uri);

  try {
    await client.connect();
    console.log('Connected to the database.');

    const databaseName = 'testdb';
    const database = client.db(databaseName);
    const collection = database.collection('data');

    const data = [
      { _id: '1', name: 'Data 1' },
      { _id: '2', name: 'Data 2' },
      { _id: '3', name: 'Data 3' },
    ];

    const result = await collection.insertMany(data);
    console.log(`${result.insertedCount} documents inserted.`);
  } catch (err) {
    console.error(err);
  } finally {
    await client.close();
    console.log('Disconnected from the database.');
  }
}

loadDatabase();
