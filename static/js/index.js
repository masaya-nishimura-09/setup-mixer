async function fetchItems(keywords) {
  if (keywords === "") return

  const keywordsList = keywords.split(",")
  console.log(keywordsList)
  let q = ""
  for (let i = 0; i < keywordsList.length; i++) {
    q += `keyword=${keywordsList[i]}`
    if (i < keywordsList.length - 1) {
      q += "&"
    }
  }

  const encodedUrl = encodeURI(`http://localhost:8080/items?${q}`);
  console.log(encodedUrl)


  try {
    const response = await fetch(encodedUrl, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });
    const items = await response.json();
    console.log(items)
    setItems(items);
  } catch (error) {
    console.error(`Internal server error.`, error);
  }
}

document.getElementById("search-button").addEventListener("click", async () => {
    const form = document.getElementById("keyword-form");
    const formData = new FormData(form);

    const keywords = formData.get("keyword")
    fetchItems(keywords);
});

const currentYear = new Date().getFullYear();
document.querySelector("#year").innerHTML = currentYear;
