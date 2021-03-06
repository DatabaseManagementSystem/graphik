<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: graphik.proto

namespace Api;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Connections is an array of Connection
 *
 * Generated from protobuf message <code>api.Connections</code>
 */
class Connections extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .api.Connection connections = 1;</code>
     */
    private $connections;
    /**
     * Generated from protobuf field <code>string seek_next = 2;</code>
     */
    private $seek_next = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Api\Connection[]|\Google\Protobuf\Internal\RepeatedField $connections
     *     @type string $seek_next
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Graphik::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .api.Connection connections = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getConnections()
    {
        return $this->connections;
    }

    /**
     * Generated from protobuf field <code>repeated .api.Connection connections = 1;</code>
     * @param \Api\Connection[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setConnections($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Api\Connection::class);
        $this->connections = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string seek_next = 2;</code>
     * @return string
     */
    public function getSeekNext()
    {
        return $this->seek_next;
    }

    /**
     * Generated from protobuf field <code>string seek_next = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setSeekNext($var)
    {
        GPBUtil::checkString($var, True);
        $this->seek_next = $var;

        return $this;
    }

}

