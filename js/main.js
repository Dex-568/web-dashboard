class DeviceTable {
    constructor(tableId) {
        this.table = document.getElementById(tableId);
        this.tbody = this.table.querySelector('tbody');
        this.init()
    }

    // not needed so much yet, maybe i'll set randomised data
    async init() {
        await this.loadData();
        setInterval(() => this.loadData(), 30000);
    }

    async loadData() {
        try {
            const response = await fetch("http://localhost:8080/devices");
            // shouldn't do this if i set it up right
            if (!response.ok) throw new Error ('HTTP error from API, status shown: ${response.status}');
            
            const devices = await response.json();
            this.updateTable(devices);
        // shouldnt do this either , but good learning
        } catch (error) {
            console.error("Unknown Fetch Error:", error);
        }
    }

    updateTable(devices) {
        this.tbody.innerHTML = '';

        devices.forEach(device => {
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${device.name}</td>
                <td>${device.status}</td>
                <td>${device.totalMemory}</td>
                <td>${device.memoryUsage}</td>
                <td>${device.memWarnThres}</td>
                <td>${device.cpuCount}</td>
                <td>${device.cpuUsage}</td>
                <td>${device.cpuWarnThres}</td>`;
            this.tbody.appendChild(row);
        });
    }

}

document.addEventListener('DOMContentLoaded', () => {
    new DeviceTable('device-table');
});