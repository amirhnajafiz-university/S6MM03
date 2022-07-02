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
        document.getElementById('contents').style.visibility = 'hidden';
        $scope.isHidden = true;
        $scope.movies2 = JSON.parse('[{"id":1,"title":"Doctor Strange in the Multiverse of Madness"},{"id":6,"title":"The Black Phone"},{"id":2,"title":"Obi-Wan Kenobi"},{"id":5,"title":"The Boys"},{"id":3,"title":"Top Gun: Maverick"},{"id":4,"title":"Stranger Things"}]')
        $scope.count = 0;
        $scope.movies2.forEach(function (value, i) {
            $scope.movies2[i].posterURL = 'localhost:8080/dash/' + value.id
        });
        $scope.myClick = function () {
            if ($scope.isHidden) {
                document.getElementById('contents').style.visibility = 'visible';
                document.getElementById('swapper').innerText = 'hide';
                $scope.isHidden = false;
            } else {
                document.getElementById('contents').style.visibility = 'hidden';
                document.getElementById('swapper').innerText = 'show';
                $scope.isHidden = true;
            }
        };
    }]
);

app.controller('secondController', ['$scope', '$routeParams', function ($scope, $routeParams) {
        $scope.isHidden = true;
        $scope.count = 0;
        $scope.movieInfo = JSON.parse('{"id":2,"title":"Obi-Wan Kenobi","director":"Deborah Chow","score":8,"description":"Jedi Master Obi-Wan Kenobi has to save young Leia after she is kidnapped, all the while being pursued by Imperial Inquisitors and his former Padawan, now known as Darth Vader.","type":"DASH1","link":""}')
        $scope.posterPath = 'localhost:8080/poster/' + $routeParams.id;
        $scope.dash = 'localhost:8080/dash/' + $routeParams.id;
        $scope.myClick = function () {
            if ($scope.isHidden) {
                document.getElementById('contents').style.visibility = 'visible';
                document.getElementById('swapper').innerText = 'hide';
                $scope.isHidden = false;
            } else {
                document.getElementById('contents').style.visibility = 'hidden';
                document.getElementById('swapper').innerText = 'show';
                $scope.isHidden = true;
            }
        };
    }]
);