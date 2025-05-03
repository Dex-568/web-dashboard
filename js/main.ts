
interface Device {
    name: string;
    status: string;
    totalMemory: string;
    memoryUsage: string;
    memWarnThres: string;
    cpuCount: string;
    cpuUsage: string;
    cpuWarnThres: string;
}

class DeviceTable {
    table: HTMLTableElement
    tbody: HTMLTableSectionElement

    constructor(tableId) {
        this.table = document.getElementById(tableId) as HTMLTableElement;
        this.tbody = this.table.querySelector('tbody') as HTMLTableSectionElement;
        this.init()
    }

    // not needed so much yet, maybe i'll set randomised data
    async init(): Promise<void> {
        await this.loadData();
        setInterval(() => this.loadData(), 30000);
    }

    async loadData(): Promise<void>{
        try {
            const response = await fetch("http://localhost:8080/devices");
            // shouldn't do this if i set it up right
            if (!response.ok) throw new Error ('HTTP error from API, status shown: ${response.status}');
            
            const devices: Device[] = await response.json();
            this.updateTable(devices);
        // shouldnt do this either , but good learning
        } catch (error) {
            console.error("Unknown Fetch Error:", error);
        }
    }

    updateTable(devices: Device[]): void {
        this.tbody.innerHTML = '';

        devices.forEach(device => {
            let exceeded = 0;
            // not de-capitalising in the html for UI experience
            // given the yes/no from the GO api, 1 exceeded is yellow
            // 2 exceeded is red

            if (device.memWarnThres.toLowerCase() === "yes") exceeded++;
            if (device.cpuWarnThres.toLowerCase() === "yes") exceeded++;

            let rowColor;

            if (device.status.toLowerCase() === "offline") {
                rowColor = "#cccccc";
            } else {
    
                if (exceeded === 0) rowColor = "green";
                    else if (exceeded === 1) rowColor = "yellow";
                    else rowColor = "red";
            }

            const row = document.createElement('tr');
            row.style.backgroundColor = rowColor
            row.innerHTML = 
                `<td>${device.name}</td>
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