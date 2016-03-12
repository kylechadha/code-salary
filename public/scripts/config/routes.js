(function() {
  angular.module('coderSalaryApp')
    .config(['$routeProvider', '$locationProvider', function($routeProvider, $locationProvider) {

      $routeProvider
        .when('/', {
          templateUrl: 'views/salary-table.html',
          controller: 'SalaryController'
        })
        .otherwise( { redirectTo: '/' } );

      $locationProvider.html5Mode(true);

    }]);
}());
