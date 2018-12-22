/**
 * @fileoverview gRPC-Web generated client stub for v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.v1 = require('./emoji_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.v1.EmojiServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.v1.EmojiServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!proto.v1.EmojiServiceClient} The delegate callback based client
   */
  this.delegateClient_ = new proto.v1.EmojiServiceClient(
      hostname, credentials, options);

};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.v1.EmojizeRequest,
 *   !proto.v1.EmojizeReply>}
 */
const methodInfo_EmojiService_Emojize = new grpc.web.AbstractClientBase.MethodInfo(
  proto.v1.EmojizeReply,
  /** @param {!proto.v1.EmojizeRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.v1.EmojizeReply.deserializeBinary
);


/**
 * @param {!proto.v1.EmojizeRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.v1.EmojizeReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.v1.EmojizeReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.v1.EmojiServiceClient.prototype.emojize =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/v1.EmojiService/Emojize',
      request,
      metadata,
      methodInfo_EmojiService_Emojize,
      callback);
};


/**
 * @param {!proto.v1.EmojizeRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.v1.EmojizeReply>}
 *     The XHR Node Readable Stream
 */
proto.v1.EmojiServicePromiseClient.prototype.emojize =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.emojize(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


module.exports = proto.v1;

