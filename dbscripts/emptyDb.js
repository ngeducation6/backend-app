// emptyDatabase.js
const { MongoClient } = require('mongodb');

async function emptyDatabase() {
//   const uri = 'mongodb://localhost:27017';
  const uri = 'mongodb://testdev:password123@localhost:27017/admin';

  const client = new MongoClient(uri);

  try {
    await client.connect();
    console.log('Connected to the database.');

    const databaseName = 'testdb';
    const database = client.db(databaseName);

    // Get all collection names in the database
    const collections = await database.listCollections().toArray();

    // Drop each collection
    for (const collection of collections) {
      await database.collection(collection.name).drop();
      console.log(`Dropped collection: ${collection.name}`);
    }

    console.log(`Emptied database: ${databaseName}`);
  } catch (err) {
    console.error(err);
  } finally {
    await client.close();
    console.log('Disconnected from the database.');
  }
}

emptyDatabase();
