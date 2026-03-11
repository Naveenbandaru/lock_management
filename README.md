# lock_management
**Predictable Lock Coordination for Scalable Multi Node Transactional Systems**

### Paper Information
- **Author(s):** Naveen Kumar Bandaru
- **Published In:** International Journal of Intelligent Systems and Applications in Engineering (IJISAE)
- **Publication Date:** Jan, 2026
- **ISSN:**  ****

### Abstract
Multi node transactional systems require efficient lock coordination to maintain consistency while supporting high concurrency. Conventional locking mechanisms rely on dynamic lock acquisition, which increases contention and coordination overhead as cluster size grows. This work introduces a deterministic lock coordination approach that enforces a predefined lock acquisition order to stabilize execution behavior. Experimental evaluation across clusters of 3 to 11 nodes shows improved throughput scalability and more predictable transaction processing performance

### Key Technical Advances in the Study
- **Deterministic Lock Coordination Framework:**  
Introduced a coordination strategy that enforces a globally consistent lock acquisition order, eliminating cyclic dependencies and reducing contention during distributed transaction execution.

- **Predictable Transaction Execution Model:**  
Developed a structured locking approach where transactions follow predefined execution paths, minimizing nondeterministic coordination delays and improving concurrency management.

- **Concurrent Transaction Processing Prototype:** 
Implemented a distributed transaction execution environment using Go based concurrency to simulate lock acquisition and evaluate throughput behavior across multiple nodes.

- **Throughput Scalability Evaluation Across Cluster Sizes:**  
Conducted experiments on clusters with 3, 5, 7, 9, and 11 nodes to analyze how deterministic lock coordination influences transaction throughput scalability.

### System Impact and Operational Advantages
- **Higher Transaction Throughput:**
Deterministic locking achieves significantly higher throughput compared with conventional locking by reducing lock contention and coordination overhead.

- **Reduced Lock Contention and Deadlocks:**  
Predefined lock acquisition order prevents cyclic dependencies and minimizes lock wait times during concurrent transaction execution.

- **Improved Scalability for Multi Node Systems:**  
Throughput continues to increase steadily as cluster size grows, demonstrating effective resource utilization and scalable transaction coordination.

- **Applicability to High Volume Transaction Platforms:**  
The framework is suitable for distributed databases, financial transaction systems, and large scale enterprise platforms requiring predictable concurrency control.

### Experimental Results (Summary)

  | Nodes | Lock Based 2PL CPU %| Lightweight Runtime Detection %| Improvment (%) |
  |-------|---------------------| -------------------------------| ---------------|
  | 3     |  88                 | 55                             | 37.50          |
  | 5     |  84                 | 49                             | 41.67          |
  | 7     |  82                 | 46                             | 43.90          |
  | 9     |  80                 | 43                             | 46.25          |
  | 11    |  79                 | 41                             | 48.10          |

### Citation
Lightweight Runtime Conflict Detection for CPU Efficient Transaction Processing
* Naveen Kumar Bandaru
* International Journal of Intelligent Systems and Applications in Engineering (IJISAE) 
* ISSN 2147-6799
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijisae.org/index.php/IJISAE \
**Author Contact** \
**LinkedIn**: linkedin.com/in/naveen-bandaru | **Email**: naveen.bandaru@gmail.com







