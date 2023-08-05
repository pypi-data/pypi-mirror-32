PyExPool is a concurrent execution pool with custom resource constraints (memory, timeouts, affinity, CPU cores and caching) and load balancing of the external applications on NUMA architecture.  All main functionality is implemented as a single-file module to be easily included into your project and customized as a part of your distribution (like in [PyCaBeM](https://github.com/eXascaleInfolab/PyCABeM)), not as a separate library. Additionally, an optional minimalistic Web interface is provided in the separate file to inspect the load balancer and execution pool. Typically, PyExPool is used as an application framework for benchmarking, load testing or other heavy-loaded multi-process execution activities on constrained computational resources.

See details on the [PyExPool page](https://github.com/eXascaleInfolab/PyExPool) and star the project if you like it! For any further assistance you can drop me a email or write [me on Linkedin](https://linkedin.com/in/artemvl).

BibTeX:
```bibtex
@misc{pyexpool,
	author = {Artem Lutov and Philippe Cudré-Mauroux},
	title = {{PyExPool-v.3: A Lightweight Execution Pool with Constraint-aware Load-Balancer.}},
	year = {2018},
	url = {https://github.com/eXascaleInfolab/PyExPool}
}
```

