const express = require("express");
const schedule = require("node-schedule");
const axios = require("axios");

const app = express();
app.use(express.json());

// API endpoint to set appointments
const WEBHOOK_URL = "http://localhost:8080/webhook"; // Replace with your actual webhook URL

function scheduleEvent(appointmentTime, message) {
  console.log(
    `Scheduling event for appointment time ${appointmentTime}: ${message}`
  );

  const job = schedule.scheduleJob(appointmentTime, async () => {
    console.log(
      `Sending webhook for appointment time ${appointmentTime}: ${message}`
    );

    try {
      const response = await axios.post(WEBHOOK_URL, {
        appointmentTime,
        message,
      });
      console.log("Webhook response:", response.data);
    } catch (error) {
      console.error("Error sending webhook:", error.message);
    }

  });
}

app.post("/schedule-appointment", (req, res) => {
  const { appointmentTime, message } = req.body;

  if (!appointmentTime || !message) {
    console.log("Invalid request received for scheduling appointment.");
    return res.status(400).json({
      error: "Invalid request. Missing appointmentTime or message.",
    });
  }

  console.log(
    `Received request to schedule appointment for ${appointmentTime}: ${message}`
  );

  scheduleEvent(appointmentTime, message);

  res.status(201).json({ message: "Appointment scheduled successfully." });
});

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
  console.log(`Scheduler server is running on port ${PORT}`);
});
