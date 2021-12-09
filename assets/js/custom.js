var MainDataTable;
$(function () {
    $('#form-data').on('submit', function (e) {
        e.preventDefault();
        ToggleLoading();
        $(this).find("button[type='submit']").prop('disabled', true);

        var formdata = null;
        if (window.FormData) {
            formdata = new FormData($('#form-data')[0]);
        }
        $.ajax({
            type: "POST",
            enctype: 'multipart/form-data',
            url: $(this).attr('action'),
            data: formdata,
            processData: false,
            contentType: false,
            success: function (respon) {
                setTimeout(() => {
                    MainDataTable.ajax.reload()
                    if (respon.status === 200) {
                        ShowFormAdd()
                        Swal.fire(
                                '',
                                respon.message,
                                'success'
                        )
                    }
                    ToggleLoading()
                }, 450)
            },
            error: function (e) {
                ToggleLoading()
            }
        });
    })
})

function CreateDataTable(config) {
    MainDataTable = $('#dataTable').DataTable({
        "processing": true,
        "serverSide": true,
        ...config
    });
}

function ShowFormAdd() {
    const form = $('#form-data')
    form.find("button[type='submit']").prop('disabled', false);
    form[0].reset();
    $('#modal-form').modal('toggle')
}
function ShowFormImport() {
    const form = $('#form-import')
    form[0].reset();
    $('#modal-import').modal('toggle')
}

function ShowFormEdit(data) {
    const form = $('#form-data')
    for (var key of Object.keys(data)) {
        form.find(".form-control[name='"+key+"']").val(data[key])
            form.find('input[type="radio"][name="'+key+'"]').eq(data[key]).prop('checked',true);

    }
    form.find("button[type='submit']").prop('disabled', false);
    $('#modal-form').modal('toggle')
}

function ToggleLoading() {
    $('#modal-loading').modal('toggle')
}

function ConfirmDelete(url, id) {
    Swal.fire({
        title: 'Are you sure?',
        text: "You won't be able to revert this!",
        icon: false,
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#d33',
        confirmButtonText: 'Delete'
    }).then((result) => {
        if (result.isConfirmed) {
            $.ajax({
                url: url,
                type: "post",
                data: {id: id},
                success: function (respon) {
                    if (respon.status === 200) {
                        Swal.fire(
                                '',
                                respon.message,
                                'success'
                        )
                    }
                    MainDataTable.ajax.reload()
                },
                error: function (jqXHR, textStatus, errorThrown) {
                }
            });
        }
    })
}

function Encode( str ) {
    return window.btoa(unescape(encodeURIComponent( str )));
}

function Decode( str ) {
    return decodeURIComponent(escape(window.atob( str )));
}
