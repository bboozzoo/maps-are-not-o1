raw <- read.csv(file="benchresults.csv",sep=",",head=TRUE)
data <- raw

png(filename="plot.png",
    units="px",
    width=640,
    height=480)

plot(data$Elements, data$Map/data$Elements, type=c("b"), log=c("x"), ylim=c(20,160),
        xlab="Number of Elements", ylab="Search Time (ns)", col="blue")
title(main="Lookup Time for Keys in a Hashmap vs. Binary Search on a List")
lines(data$Elements, y=data$Slice/data$Elements, col="red", type="b")
legend(10^2, y=160, c("Slice", "Map"), col=c("red", "blue"), lwd=c(2.5, 2.5), lty=c(1, 1))

