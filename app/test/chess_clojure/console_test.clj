(ns chess-clojure.console-test
  (:require [clojure.test :refer :all]
            [chess-clojure.console :refer [prompt]]))

(deftest console
  (testing "it should ask for keyboard input"
    (is (= (prompt) "command: "))))
