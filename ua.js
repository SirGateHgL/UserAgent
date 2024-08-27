const fs = require('fs');
const faker = require('faker');

const saveFile = process.argv[2];
const amount = parseInt(process.argv[3]);

if (!saveFile || isNaN(amount)) {
    console.log('Usage: node ua.js [save file] [amount]');
    process.exit(1);
}

const userAgents = [];

for (let i = 0; i < amount; i++) {
    userAgents.push(faker.internet.userAgent());
}

fs.writeFileSync(saveFile, userAgents.join('\n'));

console.log(`${amount} user agents saved to ${saveFile}`);
