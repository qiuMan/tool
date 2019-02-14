<form class="layui-form" action="" method="post">
  <div class="layui-form-item">
    <div class="layui-inline">
      <label class="layui-form-label">ID</label>
      <div class="layui-input-inline">
        <input type="text" name="name"  autocomplete="off" class="layui-input">
      </div>
      <button class="layui-btn" lay-submit lay-filter="search">查询</button>
    </div>
  </div>
</form>
<table class="layui-table" lay-data="{url:'/<%@#&%>/index', page:true, id:'table', method:'post', toolbar:'#hBtn'}" lay-filter="table">
    <thead>
      <tr>
<%@#tThead%>
        <th lay-data="{field:'Opt', fixed: 'right', toolbar: '#toolbar'}">操作</th>
      </tr>
    </thead>
</table>
<script type="text/html" id="hBtn">
    <div class="layui-btn-container">
        <a href="/<%@#&%>" class="layui-btn layui-btn-sm" lay-event="add">添加</a>
    </div>
</script>
  <script type="text/html" id="toolbar">
    <a href="/<%@#&%>/<%d.Id%>" class="layui-btn layui-btn-xs" lay-event="edit">编辑</a>
    <a href="<%@#&%>" class="layui-btn layui-btn-danger layui-btn-xs" lay-event="del">删除</a>
</script>
<script src="/static/js/<%@#&%>.js"></script> 

