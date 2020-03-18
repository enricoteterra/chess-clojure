(ns chess-clojure.core
  (:gen-class))

(defn -main []
  (case (read-line)
    "q" nil
    ; "a" (do (println "got a") (recur))
    ; "b" (do (println "got b") (recur))
    (do (println "Usage:\r\n q to end game\r\n") (recur))))
