const { MongoClient } = require('mongodb');

async function dropUser() {
  const uri = 'mongodb://localhost:27017';
  // const uri = 'mongodb://testdev:password123@localhost:27017/admin';
  const client = new MongoClient(uri);

  try {
    await client.connect();
    console.log('Connected to the database.');

    const adminDb = client.db('admin');
    const user = 'testdev';

    // drop User
    const dropUserResult = await adminDb.command({
      dropUser: "testdev",
    });

    console.log(`dropped user: ${user} for admin db`);

  } catch (err) {
    console.error(err);
  } finally {
    await client.close();
    console.log('Disconnected from the database.');
  }
}

dropUser();
