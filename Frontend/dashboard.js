// dashboard.js

document.addEventListener("DOMContentLoaded", () => {
    const eventSelect = document.getElementById("event");
    const metricsFields = document.getElementById("metricsFields");

    eventSelect.addEventListener("change", function () {
        // Clear the existing fields
        metricsFields.innerHTML = "";

        console.log("Changed Event!");

        // Get the selected event
        const selectedEvent = eventSelect.value;

        if (selectedEvent) {
            // Dynamically create fields based on the selected event
            const eventLabel = document.createElement("label");
            eventLabel.textContent = selectedEvent.replace(/_/g, " ").toUpperCase();
            metricsFields.appendChild(eventLabel);

            metricsFields.appendChild(document.createElement("br"));

            // Create fields for time, confidence, intelligence, and energy
            const timeInput = document.createElement("input");
            timeInput.type = "number";
            timeInput.id = `${selectedEvent}_time`;
            timeInput.name = `${selectedEvent}_time`;
            timeInput.placeholder = "Time";
            metricsFields.appendChild(timeInput);

            metricsFields.appendChild(document.createElement("br"));

            const confidenceInput = document.createElement("input");
            confidenceInput.type = "number";
            confidenceInput.id = `${selectedEvent}_confidence`;
            confidenceInput.name = `${selectedEvent}_confidence`;
            confidenceInput.placeholder = "Confidence";
            metricsFields.appendChild(confidenceInput);

            metricsFields.appendChild(document.createElement("br"));

            const intelligenceInput = document.createElement("input");
            intelligenceInput.type = "number";
            intelligenceInput.id = `${selectedEvent}_intelligence`;
            intelligenceInput.name = `${selectedEvent}_intelligence`;
            intelligenceInput.placeholder = "Intelligence";
            metricsFields.appendChild(intelligenceInput);

            metricsFields.appendChild(document.createElement("br"));

            const energyInput = document.createElement("input");
            energyInput.type = "number";
            energyInput.id = `${selectedEvent}_energy`;
            energyInput.name = `${selectedEvent}_energy`;
            energyInput.placeholder = "Energy";
            metricsFields.appendChild(energyInput);

            metricsFields.appendChild(document.createElement("br"));
        }
    });

    document.getElementById("dashboardForm").addEventListener("submit", async function (event) {
        event.preventDefault();


        console.log("Sent data to table!");
        const selectedEvent = document.getElementById("event").value;  // Get selected event

        // Only proceed if an event is selected
        if (!selectedEvent) {
            document.getElementById("message").innerText = "Please select an event.";
            return;
        }

        const data = {
            username: "user1", // Replace with logged-in user's username
            selectedEvent: selectedEvent,  // Selected event from dropdown
            time: document.getElementById(`${selectedEvent}_time`).value,
            confidence: document.getElementById(`${selectedEvent}_confidence`).value,
            intelligence: document.getElementById(`${selectedEvent}_intelligence`).value,
            energy: document.getElementById(`${selectedEvent}_energy`).value
        };

        // Check if all fields are filled before submitting
        if (!data.time || !data.confidence || !data.intelligence || !data.energy) {
            document.getElementById("message").innerText = "Please fill all fields.";
            return;
        }

        console.log("Sending data");

        const response = await fetch("/submitData", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(data)
        });

        console.log("done sending");

        const result = await response.json();
        document.getElementById("message").innerText = result.message;
    });
});
