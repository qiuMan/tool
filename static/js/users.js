
layui.use(['form', 'table', 'jquery'], function(){
    var form = layui.form
    ,table = layui.table
    ,$ = layui.jquery;

     //监听提交
  form.on('submit(search)', function(data){
      table.reload('table', {
        where: data.field //设定异步数据接口的额外参数
      });
      return false;
  });

  form.on('submit(save)', function(data){
    var method = "POST"
    if (data.field.Id > 0) {
        method = "PUT"
    } 
    $.ajax({
         url: '',
         type: 'json',
         method: method,
         data: data.field,
         success: function (ret) {
            if (ret.code == 0) {
                layer.msg("操作成功", {icon: 1, time: 1000}, function () {
                    
                })
            } else {
                layer.msg(ret.msg, {icon: 2})
            }
         }
    })
    return false;
  });


  table.on("toolbar(table)", function (obj) {
      var layEevent = obj.event;
      
      if (layEevent == "add") {
        
      }
  })

  table.on("tool(table)", function (obj) {
        var data = obj.data;
        var layEevent = obj.event;
        if (layEevent == 'edit') {
            
        } else if (layEevent == 'del') {
            layer.confirm('确定删除？', {icon: 3, title:'提示'}, function(index){
                $.ajax({
                    url: '/users/' + data.Id,
                    type: 'json',
                    method: 'Delete',
                    success: function (ret) {
                       if (ret.code == 0) {
                           layer.msg("操作成功", {icon: 1, time: 1000}, function () {
                            layer.close(index);
                             window.location.reload()
                           })
                       } else {
                           layer.msg(ret.msg, {icon: 2})
                       }
                    }
                })
                
              });
           
        }
  })
})
