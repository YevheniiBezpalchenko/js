const query = async () => {
    let res = await fetch('http://localhost:8080/get_products');
    let products = await res.json();
    let flag = document.getElementById("flag");
    for(let product of products){
        let product_div = document.createElement("div");
        product_div.innerHTML = `${product.name} - ${product.price}`;
        flag.appendChild(product_div);
    }
}
query()