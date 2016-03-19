(function() {
	var SalaryController = function($scope, $http) {
		$scope.salaries = [];
		$scope.stack = [];
		$scope.sortType;
		$scope.sortReverse;

		// When landing on the page, get all the salaries and show them.
		$http.get('/api/salaryData')
			.success(function(data) {
				$scope.salaries = data;
				console.log(data);
			})
			.error(function(data) {
				console.log("Error: " + data);
			})

			// * Will need to add validation here
			$scope.createSalary = function() {
				$scope.formData.stack = [];
				if ($scope.stack[0]) {
					$scope.formData.stack.push("Ruby on Rails")
				}
				if ($scope.stack[1]) {
					$scope.formData.stack.push("Node.js")
				}
				if ($scope.stack[2]) {
					$scope.formData.stack.push("Python")
				}
				console.log($scope.formData);

				$http.post('/api/salaryData', $scope.formData)
					.success(function(data) {
						$scope.formData = {};
						$scope.salaries = data;
						console.log(data);
					})
					.error(function(data) {
						console.log("Error: " + data);
					})
			}
	};

	SalaryController.$inject = ['$scope', '$http'];
	angular.module('coderSalaryApp').controller('SalaryController', SalaryController);
}());
