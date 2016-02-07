(function() {

  angular.module('omniaApp')
    .config(['$routeProvider', '$locationProvider', function($routeProvider, $locationProvider) {

      $routeProvider
        .when('/', {})
        .when('/users', {
          templateUrl: 'views/users-create.html',
          controller: 'UserController'
        })
        .otherwise( { redirectTo: '/' } );

      $locationProvider.html5Mode(true);

    }]);

}());
