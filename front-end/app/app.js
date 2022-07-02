var app = angular.module('myApp', ['ngRoute'])


app.config(function ($routeProvider) {
    $routeProvider
        .when('/', {
            templateUrl: 'home/first.html',
            controller: 'myCtrl'
        })
        .when('/movies/:id', {
            templateUrl: 'home/second.html',
            controller: 'secondController'
        })
});

app.controller('myCtrl', ['$scope', function ($scope) {

        $scope.movies2 = JSON.parse('[{"id":1,"title":"Doctor Strange in the Multiverse of Madness"},{"id":6,"title":"The Black Phone"},{"id":2,"title":"Obi-Wan Kenobi"},{"id":5,"title":"The Boys"},{"id":3,"title":"Top Gun: Maverick"},{"id":4,"title":"Stranger Things"}]')
        $scope.count = 0;
        $scope.myClick = function () {
            $scope.count++;
        };
    }]
);

app.controller('secondController', ['$scope', '$routeParams', function ($scope, $routeParams) {

        $scope.count = 0;
        $scope.movieInfo = JSON.parse('{"id":2,"title":"Obi-Wan Kenobi","director":"Deborah Chow","score":8,"description":"Jedi Master Obi-Wan Kenobi has to save young Leia after she is kidnapped, all the while being pursued by Imperial Inquisitors and his former Padawan, now known as Darth Vader.","type":"DASH1","link":""}')
        $scope.posterPath = 'localhost:8080/poster/' + $routeParams.id;
        $scope.dash = 'localhost:8080/dash/' + $routeParams.id;
        $scope.myClick = function () {
            $scope.count++;
        }
    }]
);