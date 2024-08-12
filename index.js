const categories = [
  { name: "Transportation" },
  { name: "Home Energy" },
  { name: "Food" },
  { name: "Waste" },
  { name: "Water Usage" },
];

const activities = [
  {
    name: "Public transport (bus)",
    carbonFootprint: 0.1,
    unit: "km",
    category_id: "b027d780-f97c-4fb4-aac0-4e05357198ef",
  },
  {
    name: "Public transport (train)",
    carbonFootprint: 0.04,
    unit: "km",
    category_id: "b027d780-f97c-4fb4-aac0-4e05357198ef",
  },
  {
    name: "Private car (petrol)",
    carbonFootprint: 0.192,
    unit: "km",
    category_id: "b027d780-f97c-4fb4-aac0-4e05357198ef",
  },
  {
    name: "Private car (diesel)",
    carbonFootprint: 0.171,
    unit: "km",
    category_id: "b027d780-f97c-4fb4-aac0-4e05357198ef",
  },
  {
    name: "Electric car",
    carbonFootprint: 0.053,
    unit: "km",
    category_id: "b027d780-f97c-4fb4-aac0-4e05357198ef",
  },
  {
    name: "Motorcycle",
    carbonFootprint: 0.103,
    unit: "km",
    category_id: "b027d780-f97c-4fb4-aac0-4e05357198ef",
  },
  {
    name: "Bicycle",
    carbonFootprint: 0,
    unit: "km",
    category_id: "b027d780-f97c-4fb4-aac0-4e05357198ef",
  },
  {
    name: "Walking",
    carbonFootprint: 0,
    unit: "km",
    category_id: "b027d780-f97c-4fb4-aac0-4e05357198ef",
  },
  {
    name: "Electricity consumption",
    carbonFootprint: 0.475,
    unit: "kWh",
    category_id: "41d6f059-fdac-40d9-8fa1-302c6babe91c",
  },
  {
    name: "Natural gas",
    carbonFootprint: 0.203,
    unit: "kWh",
    category_id: "41d6f059-fdac-40d9-8fa1-302c6babe91c",
  },
  {
    name: "Heating oil",
    carbonFootprint: 0.247,
    unit: "kWh",
    category_id: "41d6f059-fdac-40d9-8fa1-302c6babe91c",
  },
  {
    name: "Beef",
    carbonFootprint: 27,
    unit: "kg",
    category_id: "49981b4a-3fb6-45bb-8923-a454220d53a9",
  },
  {
    name: "Lamb",
    carbonFootprint: 39.2,
    unit: "kg",
    category_id: "49981b4a-3fb6-45bb-8923-a454220d53a9",
  },
  {
    name: "Pork",
    carbonFootprint: 12.1,
    unit: "kg",
    category_id: "49981b4a-3fb6-45bb-8923-a454220d53a9",
  },
  {
    name: "Chicken",
    carbonFootprint: 6.9,
    unit: "kg",
    category_id: "49981b4a-3fb6-45bb-8923-a454220d53a9",
  },
  {
    name: "Fish",
    carbonFootprint: 6.1,
    unit: "kg",
    category_id: "49981b4a-3fb6-45bb-8923-a454220d53a9",
  },
  {
    name: "Eggs",
    carbonFootprint: 4.8,
    unit: "kg",
    category_id: "49981b4a-3fb6-45bb-8923-a454220d53a9",
  },
  {
    name: "Rice",
    carbonFootprint: 2.7,
    unit: "kg",
    category_id: "49981b4a-3fb6-45bb-8923-a454220d53a9",
  },
  {
    name: "Potatoes",
    carbonFootprint: 0.3,
    unit: "kg",
    category_id: "49981b4a-3fb6-45bb-8923-a454220d53a9",
  },
  {
    name: "Vegetables (average)",
    carbonFootprint: 2,
    unit: "kg",
    category_id: "49981b4a-3fb6-45bb-8923-a454220d53a9",
  },
  {
    name: "Fruits (average)",
    carbonFootprint: 1.1,
    unit: "kg",
    category_id: "49981b4a-3fb6-45bb-8923-a454220d53a9",
  },
  {
    name: "Landfill waste",
    carbonFootprint: 0.587,
    unit: "kg",
    category_id: "2ea8615b-3d9d-45a9-add1-6056e92d2adb",
  },
  {
    name: "Recycled waste",
    carbonFootprint: 0.021,
    unit: "kg",
    category_id: "2ea8615b-3d9d-45a9-add1-6056e92d2adb",
  },
  {
    name: "Composted waste",
    carbonFootprint: 0.008,
    unit: "kg",
    category_id: "2ea8615b-3d9d-45a9-add1-6056e92d2adb",
  },
  {
    name: "Water consumption",
    carbonFootprint: 0.344,
    unit: "m³",
    category_id: "9e44f831-43b3-4a56-a6cd-0b7f2fb1ee5e",
  },
];

const API_URL = "http://localhost:3000/";

async function insertCategories(categories) {
  for (const category of categories) {
    try {
      const response = await fetch(`http://localhost:3000/category/create`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(category),
      });
      const data = await response.json();
      console.log("Category Response:", data);
    } catch (error) {
      console.error("Error inserting category:", category.name, error);
    }
  }
}

async function insertActivities(activities) {
  for (const activity of activities) {
    await fetch(`http://localhost:3000/activity/create`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(activity),
    });
  }
}

async function main() {
  //   await insertCategories(categories);
  await insertActivities(activities);
}

main().catch((error) => console.error("Error:", error));
