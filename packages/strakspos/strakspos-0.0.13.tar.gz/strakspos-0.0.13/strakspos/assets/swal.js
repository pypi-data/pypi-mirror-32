function swal_it(title, msg, icon)
{
    swal({
        title:title, 
        content: {
            element: "p",
            attributes: {
                innerHTML: '<p style="padding-bottom:1px;">' + msg + '</p>'
            },
        },
        icon:icon,
        buttons: false,
        timer:8000,
    });
}
