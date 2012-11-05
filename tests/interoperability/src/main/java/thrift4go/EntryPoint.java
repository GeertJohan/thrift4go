package thrift4go;

import org.apache.thrift.transport.TTransportException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import thrift4go.enumm.EnummServerEntryPoint;

public class EntryPoint {
	private static final Logger log = LoggerFactory.getLogger(EntryPoint.class);

	public static void main(final String[] args) {

		try {
			// Start enumm test servers
			new Thread() {
				public void run() {
					try {
						EnummServerEntryPoint.start("BINARY", 8080);
					} catch(TTransportException tte) {
						tte.printStackTrace();
					}
				}
			}.start();

			new Thread() {
				public void run() {
					try {
						EnummServerEntryPoint.start("JSON", 8081);
					} catch(TTransportException tte) {
						tte.printStackTrace();
					}
				}
			}.start();

			new Thread() {
				public void run() {
					try {
						EnummServerEntryPoint.start("COMPACT", 8082);
					} catch(TTransportException tte) {
						tte.printStackTrace();
					}
				}
			}.start();

			// Start listt test servers
			new Thread() {
				public void run() {
					// ListtServerEntryPoint.start("BINARY", 8180);
				}
			}.start();

			new Thread() {
				public void run() {
					// ListtServerEntryPoint.start("JSON", 8181);
				}
			}.start();

			new Thread() {
				public void run() {
					// ListtServerEntryPoint.start("COMPACT", 8182);
				}
			}.start();
			
		} catch(Exception e) {
			e.printStackTrace();
		}
	}
}