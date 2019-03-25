let FORM_PRODUCT = $('#form_product');

//MODEL
let product = {};

FORM_PRODUCT.on('submit', (e) => {
    e.preventDefault();
    product.name = $('#name').val();
    $.ajax({
        url: 'http://localhost:93/api/product',
        type: 'POST',
        data: JSON.stringify(product),
        contentType: "application/json; charset=utf-8",
        dataType:"json",
        success: (resp) => {
            alert(resp.Mensaje)
            console.log(resp);
        }
    })
});