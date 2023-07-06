const { MongoClient } = require('mongodb');

async function createUser() {
  const uri = 'mongodb://localhost:27017';
  // const uri = 'mongodb://testdev:password123@localhost:27017/admin';
  const client = new MongoClient(uri);

  try {
    await client.connect();
    console.log('Connected to the database.');

    const databaseName = 'testdb';
    const adminDb = client.db('admin');


    // create a user with a hardcoded password
    const user = 'testdev';


    const createUserResult = await adminDb.command({
      createUser: "testdev",
      pwd: "password123",
      roles: [{ role: "userAdminAnyDatabase", db: "admin" }]
    });
    console.log(`Created user: ${user} for admin db`);

  } catch (err) {
    console.error(err);
  } finally {
    await client.close();
    console.log('Disconnected from the database.');
  }
}

createUser();
