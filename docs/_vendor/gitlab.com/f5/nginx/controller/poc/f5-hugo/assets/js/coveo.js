document.addEventListener('DOMContentLoaded', function () {
  // 2. Configure a search endpoint
  Coveo.SearchEndpoint.configureCloudV2Endpoint("", 'xxfabb573f-7116-4a17-ac25-b18118218ae0');
  const root = document.getElementById("search");

  const searchBoxRoot = document.getElementById("searchbox");
  Coveo.initSearchbox(
    searchBoxRoot,
    "/search.html");

  var resetbtn = document.querySelector('#reset_btn');
  resetbtn.onclick = function () {
    document.querySelector('.coveo-facet-header-eraser').click();
  };
  Coveo.$$(root).on("querySuccess", function(e, args) {
      resetbtn.style.display="block";
    });

  Coveo.init(root);
})

