// dropDatabase.js
const { MongoClient } = require('mongodb');

async function dropDatabase() {
  const uri = 'mongodb://localhost:27017';
  const client = new MongoClient(uri);

  try {
    await client.connect();
    console.log('Connected to the database.');

    const databaseName = 'testdb';
    const database = client.db(databaseName);

    // Drop the database
    await database.dropDatabase();
    console.log(`Dropped database: ${databaseName}`);
  } catch (err) {
    console.error(err);
  } finally {
    await client.close();
    console.log('Disconnected from the database.');
  }
}

dropDatabase();
