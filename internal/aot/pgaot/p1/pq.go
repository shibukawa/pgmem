package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pq_cleanup_redirect_to_shm_mq(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_pq_cleanup_redirect_to_shm_mq[0])) = v4
	*(*int32)(unsafe.Add(mBase, _c_F_pq_cleanup_redirect_to_shm_mq[1])) = v4
	return
}
func F_pq_getmsgbyte(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v4 <= v3 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16908800))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_pq_getmsgbyte_0), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pq_getmsgbyte_1), int32(404), int32(_a_F_pq_getmsgbyte_2))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v3 + int32(1)
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v3))))
		return v29
	}
}
func F_pq_gettcpusertimeout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 == v2 {
		v21 = v2
	} else {
		v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
		if v10 == int32(1) {
			v21 = v2
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+412))
			if v13 != 0 {
				v21 = v13
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
				if v14 != 0 {
					v21 = v14
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(4)
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(396))))
					v21 = v20
				}
			}
		}
	}
	m.G0 = v6 + int32(16)
	return v21
}
func F_pq_recvbuf(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pq_recvbuf[0]))
	if int32(0) < v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_pq_recvbuf[1]))
	if v4 < v8 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_pq_recvbuf[2]))
	if v28 != 0 {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(0)
	goto L3
L5:
	;
	v10 = v8 - v4
	if v10 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pq_recvbuf[0])) = int32(0)
	v23 = int32(_a_F_pq_recvbuf_0)
	goto L4
L8:
	;
	v11 = int32(_a_F_pq_recvbuf_1)
	base.MemoryCopy(m, v11, v4+v11, v10)
	goto L10
L9:
	;
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pq_recvbuf[1])) = v10
	v23 = int32(_a_F_pq_recvbuf_2)
	goto L4
L11:
	;
	v29 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+4)) = uint8(v29)
	v31 = int32(-1)
	goto L16
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L18
	} else {
		goto L29
	}
L14:
	;
	return v85
L15:
	;
	if v45 == int32(0) {
		v85 = v31
		goto L14
	} else {
		goto L28
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pq_recvbuf[3])) = int32(0)
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_pq_recvbuf[2]))
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_pq_recvbuf[1]))
	v45 = F_secure_read(m, v38, v40+int32(_a_F_pq_recvbuf_1), int32(_a_F_pq_recvbuf_3)-v40)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v52 == int32(0) {
		v85 = v31
		goto L14
	} else {
		goto L22
	}
L18:
	;
	return int32(0)
L19:
	;
	if int32(0) <= v45 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_pq_recvbuf[3]))
	if v52 == int32(27) {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	v59 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	if v59 == int32(0) {
		v85 = v31
		goto L14
	} else {
		goto L24
	}
L24:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	F_errmsg(m, int32(_a_F_pq_recvbuf_4), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_pq_recvbuf_5), int32(942), int32(_a_F_pq_recvbuf_6))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	return int32(-1)
L28:
	;
	v79 = int32(_a_F_pq_recvbuf_0)
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_pq_recvbuf[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pq_recvbuf[1])) = v81 + v45
	v85 = int32(0)
	goto L14
L29:
	;
	F_errcode(m, int32(50332160))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L18
	} else {
		goto L30
	}
L30:
	;
	F_errmsg(m, int32(_a_F_pq_recvbuf_7), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L18
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_pq_recvbuf_5), int32(886), int32(_a_F_pq_recvbuf_8))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L18
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pq_sendcountedtext(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v52 int32
	_ = v52
	v5 = F_pg_server_to_client(m, l1, l2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		if l1 != v5 {
			v8 = F_strlen(m, v5)
			mBase = m.M
			F_enlargeStringInfo(m, l0, int32(4))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v17 = int32(16711935)
				*(*int32)(unsafe.Add(mBase, uint32(v12+v13))) = base.I32_rotr(v8, int32(24))&v17 | base.I32_rotr(v8&v17, int32(8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12 + int32(4)
				F_appendBinaryStringInfoNT(m, l0, v5, v8)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_pfree(m, v5)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			F_enlargeStringInfo(m, l0, int32(4))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v40 = int32(16711935)
				*(*int32)(unsafe.Add(mBase, uint32(v35+v36))) = base.I32_rotr(l2, int32(24))&v40 | base.I32_rotr(l2&v40, int32(8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v35 + int32(4)
				F_appendBinaryStringInfoNT(m, l0, l1, l2)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_pq_sendtext(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v5 = F_pg_server_to_client(m, l1, l2)
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		if l1 != v5 {
			v8 = F_strlen(m, v5)
			F_appendBinaryStringInfo(m, l0, v5, v8)
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				F_pfree(m, v5)
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			F_appendBinaryStringInfo(m, l0, l1, l2)
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_pq_setkeepalivesidle(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
	if l1 == int32(0) {
	} else {
		v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
		if v11 == int32(1) {
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+400))
			if l0 == v14 {
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+384))
				if int32(0) < v16 {
					if l0 == int32(0) {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+384))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v46
					} else {
					}
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+400)) = v49
				} else {
					v19 = int32(0)
					v21 = m.G0
					v23 = v21 - int32(16)
					m.G0 = v23
					if l1 == v19 {
						v38 = v19
					} else {
						v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
						if v27 == int32(1) {
							v38 = v19
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+400))
							if v30 != 0 {
								v38 = v30
							} else {
								v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+384))
								if v31 != 0 {
									v38 = v31
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(4)
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(384))))
									v38 = v37
								}
							}
						}
					}
					m.G0 = v23 + int32(16)
					if int32(0) <= v38 {
						if l0 == int32(0) {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+384))
							*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v46
						} else {
						}
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+400)) = v49
					} else {
					}
				}
			}
		}
	}
	m.G0 = v6 + int32(16)
	return
}
