package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_logicalrep_get_attrs_str(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	F_initStringInfo(m, v8+int32(16))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l1 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	if int32(0) <= v72 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v72 = base.I32_ctz(v58) | v59<<(uint(int32(5))%32)
	goto L3
L5:
	;
	v72 = int32(-2)
	goto L3
L6:
	;
	v25 = base.I32_div_s(int32(0), int32(32))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v26 <= v25 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v29 = l1 + int32(8)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v25<<(uint(int32(2))%32))))
	v36 = v33 & int32(-1)
	if v36 != 0 {
		v58 = v36
		v59 = v25
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v38 = v25 + int32(1)
	if v38 == v26 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v41 = v38
	goto L10
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v29+v41<<(uint(int32(2))%32))))
	if v48 != 0 {
		v58 = v48
		v59 = v41
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L5
L12:
	;
	v50 = v41 + int32(1)
	if v50 != v26 {
		v41 = v50
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v78 = v72
	v79 = int32(0)
	goto L17
L15:
	;
	goto L16
L16:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	m.G0 = v8 + int32(32)
	return v163
L17:
	;
	v81 = v79 + int32(1)
	if int32(2) <= v81 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L16
L19:
	;
	F_appendStringInfoString(m, v8+int32(16), int32(704244))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v78<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v93
	F_appendStringInfo(m, v8+int32(16), int32(686384), v8)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	if l1 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	if int32(0) <= v155 {
		v78 = v155
		v79 = v81
		goto L17
	} else {
		goto L35
	}
L25:
	;
	v155 = base.I32_ctz(v141) | v142<<(uint(int32(5))%32)
	goto L24
L26:
	;
	v155 = int32(-2)
	goto L24
L27:
	;
	v106 = v78 + int32(1)
	v108 = base.I32_div_s(v106, int32(32))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v109 <= v108 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v112 = l1 + int32(8)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112+v108<<(uint(int32(2))%32))))
	v119 = v116 & (int32(-1) << (uint(v106) % 32))
	if v119 != 0 {
		v141 = v119
		v142 = v108
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v121 = v108 + int32(1)
	if v121 == v109 {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v124 = v121
	goto L31
L31:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v112+v124<<(uint(int32(2))%32))))
	if v131 != 0 {
		v141 = v131
		v142 = v124
		goto L25
	} else {
		goto L33
	}
L32:
	;
	goto L26
L33:
	;
	v133 = v124 + int32(1)
	if v133 != v109 {
		v124 = v133
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	goto L18
}
func F_logicalrep_launcher_attach_dshmem(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	v4 = *(*int32)(unsafe.Add(mBase, _consts[514]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	if v5 != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, _consts[515]))
		if v7 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, _consts[29]))
			v13 = F_LWLockAcquire(m, v9+int32(5504), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				v15 = int32(4442992)
				v16 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				v19 = *(*int32)(unsafe.Add(mBase, _consts[87]))
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v19
				v22 = *(*int32)(unsafe.Add(mBase, _consts[514]))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
				if v23 == int32(0) {
					v30 = F_dsa_create_ext(m, int32(82), int32(1048576), int32(134217728))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[516])) = v30
						F_dsa_pin(m, v30)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, _consts[516]))
							F_dsa_pin_mapping(m, v36)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, _consts[516]))
								v44 = F_dshash_create(m, v41, int32(1579080), int32(0))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[515])) = v44
									v48 = *(*int32)(unsafe.Add(mBase, _consts[516]))
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+28))
									v52 = *(*int32)(unsafe.Add(mBase, _consts[514]))
									*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v50
									v55 = *(*int32)(unsafe.Add(mBase, _consts[515]))
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+32))
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
									v59 = *(*int32)(unsafe.Add(mBase, _consts[514]))
									*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v57
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
									v85 = *(*int32)(unsafe.Add(mBase, _consts[29]))
									F_LWLockRelease(m, v85+int32(5504))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				} else {
					v62 = *(*int32)(unsafe.Add(mBase, _consts[515]))
					if v62 != 0 {
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
						v85 = *(*int32)(unsafe.Add(mBase, _consts[29]))
						F_LWLockRelease(m, v85+int32(5504))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return
						} else {
							return
						}
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
						v65 = F_dsa_attach(m, v64)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[516])) = v65
							F_dsa_pin_mapping(m, v65)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								v72 = *(*int32)(unsafe.Add(mBase, _consts[516]))
								v75 = *(*int32)(unsafe.Add(mBase, _consts[514]))
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
								v78 = F_dshash_attach(m, v72, int32(1579080), v76, int32(0))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[515])) = v78
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
									v85 = *(*int32)(unsafe.Add(mBase, _consts[29]))
									F_LWLockRelease(m, v85+int32(5504))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[29]))
		v13 = F_LWLockAcquire(m, v9+int32(5504), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v15 = int32(4442992)
			v16 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v19 = *(*int32)(unsafe.Add(mBase, _consts[87]))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v19
			v22 = *(*int32)(unsafe.Add(mBase, _consts[514]))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
			if v23 == int32(0) {
				v30 = F_dsa_create_ext(m, int32(82), int32(1048576), int32(134217728))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[516])) = v30
					F_dsa_pin(m, v30)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, _consts[516]))
						F_dsa_pin_mapping(m, v36)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, _consts[516]))
							v44 = F_dshash_create(m, v41, int32(1579080), int32(0))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[515])) = v44
								v48 = *(*int32)(unsafe.Add(mBase, _consts[516]))
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+28))
								v52 = *(*int32)(unsafe.Add(mBase, _consts[514]))
								*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v50
								v55 = *(*int32)(unsafe.Add(mBase, _consts[515]))
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+32))
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
								v59 = *(*int32)(unsafe.Add(mBase, _consts[514]))
								*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v57
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
								v85 = *(*int32)(unsafe.Add(mBase, _consts[29]))
								F_LWLockRelease(m, v85+int32(5504))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			} else {
				v62 = *(*int32)(unsafe.Add(mBase, _consts[515]))
				if v62 != 0 {
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
					v85 = *(*int32)(unsafe.Add(mBase, _consts[29]))
					F_LWLockRelease(m, v85+int32(5504))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return
					} else {
						return
					}
				} else {
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
					v65 = F_dsa_attach(m, v64)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[516])) = v65
						F_dsa_pin_mapping(m, v65)
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							v72 = *(*int32)(unsafe.Add(mBase, _consts[516]))
							v75 = *(*int32)(unsafe.Add(mBase, _consts[514]))
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
							v78 = F_dshash_attach(m, v72, int32(1579080), v76, int32(0))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[515])) = v78
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
								v85 = *(*int32)(unsafe.Add(mBase, _consts[29]))
								F_LWLockRelease(m, v85+int32(5504))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_logicalrep_read_tuple(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v85 int32
	_ = v85
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = F_pq_getmsgint(m, l0, int32(2))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = F_palloc0(m, v15<<(uint(int32(4))%32))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v19
	v22 = F_palloc(m, v15)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v22
	if int32(0) < v15 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v32 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	m.G0 = v12 + int32(16)
	return
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v39 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L7
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v41+v32))) = uint8(v39)
	v44 = base.I32_extend8_s(v39)
	switch v44 - int32(98) {
	case 0, 18:
		goto L12
	default:
		goto L13
	case 12, 19:
		goto L11
	}
L11:
	;
	v85 = v32 + int32(1)
	if v85 != v15 {
		v32 = v85
		goto L8
	} else {
		goto L20
	}
L12:
	;
	v61 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L17
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v44
	F_errmsg_internal(m, int32(648062), v12)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(473147), int32(912), int32(364960))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	v64 = v61 + int32(1)
	v65 = F_palloc(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_pq_copymsgbytes(m, l0, v65, v61)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v70 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v61+v65))) = uint8(v70)
	v74 = v38 + v32<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v74)+12)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v65
	goto L11
L20:
	;
	goto L9
}
func F_logicalrep_relmap_invalidate_cb(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[526]))
	if v9 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v6 + int32(32)
	return
L2:
	;
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_hash_seq_init(m, v6+int32(12), v9)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_hash_seq_init(m, v6+int32(12), v9)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L14
	}
L6:
	;
	return
L7:
	;
	goto L8
L8:
	;
	v21 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v27 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+32)) = uint8(v27)
	F_hash_seq_term(m, v6+int32(12))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L13
	}
L10:
	;
	if v21 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	if v25 != l1 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	goto L1
L14:
	;
	v39 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	if v39 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v43 = v39
	goto L17
L17:
	;
	v46 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+32)) = uint8(v46)
	v50 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L19
	}
L18:
	;
	goto L1
L19:
	;
	if v50 != 0 {
		v43 = v50
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
}
func F_logicalrep_relmap_update(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[526]))
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v45 = v13
	goto L3
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _consts[527]))
	if v15 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v47 = F_hash_search(m, v45, l0, int32(1), v10)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L7
	} else {
		goto L11
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	v25 = F_AllocSetContextCreateInternal(m, v20, int32(59501), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v28 = v15
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(309237645316)
	v36 = F_hash_create(m, int32(380345), int32(128), v10, int32(1064))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[527])) = v25
	v28 = v25
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[526])) = v36
	F_CacheRegisterRelcacheCallback(m, int32(1014))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[526]))
	v45 = v43
	goto L3
L11:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v49 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_logicalrep_relmap_free_entry(m, v47)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v57 = F__emscripten_memset_bulkmem(m, v47, base.I32_extend8_s(int32(0)), int32(72))
	mBase = m.M
	goto L16
L15:
	;
	goto L14
L16:
	;
	v58 = int32(4442992)
	v59 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v62 = *(*int32)(unsafe.Add(mBase, _consts[527]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v64
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v67 = F_pstrdup(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v67
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v71 = F_pstrdup(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v71
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = v74
	v78 = F_palloc(m, v74<<(uint(int32(2))%32))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+16)) = v78
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v84 = F_palloc(m, v81<<(uint(int32(2))%32))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v84
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v87 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v94 = int32(0)
	goto L24
L22:
	;
	goto L23
L23:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+24)) = uint8(v124)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v127 = F_bms_copy(m, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L7
	} else {
		goto L28
	}
L24:
	;
	v98 = v94 << (uint(int32(2)) % 32)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v98+v99)))
	v102 = F_pstrdup(m, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L7
	} else {
		goto L26
	}
L25:
	;
	goto L23
L26:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v104+v98))) = v102
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v109+v98)))
	*(*int32)(unsafe.Add(mBase, uint32(v107+v98))) = v111
	v114 = v94 + int32(1)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v114 < v115 {
		v94 = v114
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+28)) = v127
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v59
	m.G0 = v10 + int32(48)
	return
}
func F_logicalrep_worker_stop(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	v8 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v12 = F_LWLockAcquire(m, v8+int32(5504), int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _consts[513]))
	if v15 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v54+int32(5504))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L15
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[514]))
	v25 = int32(0)
	goto L5
L5:
	;
	v30 = v19 + int32(16) + v25*int32(112)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+16)))
	if v31 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	F_logicalrep_worker_stop_internal(m, v30, int32(15))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L14
	}
L7:
	;
	goto L6
L8:
	;
	v42 = v25 + int32(1)
	if v42 != v15 {
		v25 = v42
		goto L5
	} else {
		goto L13
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v34 == int32(3) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+32))
	if v37 != l0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v30)+36))
	if v39 == l1 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	goto L3
L14:
	;
	goto L3
L15:
	;
	return
}
func F_logicalrep_worker_stop_internal(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v7&int32(1) == int32(0) {
		v100 = v5
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v100)+44))
	v103 = F_kill(m, v102, l1)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L5
	} else {
		goto L30
	}
L3:
	;
	if v5 != 0 {
		v100 = v5
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v13+int32(5504))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	v23 = F_WaitLatch(m, v19, int32(41), int32(10), int32(134217734))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L8
	}
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v44 = F_LWLockAcquire(m, v40+int32(5504), int32(1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L13
	}
L8:
	;
	if v23&int32(1) == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(0)
	goto L10
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v34 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	goto L7
L13:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v46 != int32(1) {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v49 != v6 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v51 != 0 {
		v100 = v51
		goto L2
	} else {
		goto L16
	}
L16:
	;
	goto L17
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v57+int32(5504))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L19
	}
L18:
	;
	v100 = v95
	goto L2
L19:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	v67 = F_WaitLatch(m, v63, int32(41), int32(10), int32(134217734))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L5
	} else {
		goto L21
	}
L20:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v88 = F_LWLockAcquire(m, v84+int32(5504), int32(1))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L26
	}
L21:
	;
	if v67&int32(1) == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
	goto L23
L23:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v78 == int32(0) {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	goto L20
L26:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v90 != int32(1) {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v93 != v6 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v95 == int32(0) {
		goto L17
	} else {
		goto L29
	}
L29:
	;
	goto L18
L30:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v105 == int32(0) {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L32
L32:
	;
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v112 != v6 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	goto L1
L34:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v115+int32(5504))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	v125 = F_WaitLatch(m, v121, int32(41), int32(10), int32(134217733))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L37
	}
L36:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v146 = F_LWLockAcquire(m, v142+int32(5504), int32(1))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L5
	} else {
		goto L42
	}
L37:
	;
	if v125&int32(1) == int32(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = int32(0)
	goto L39
L39:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v136 == int32(0) {
		goto L36
	} else {
		goto L40
	}
L40:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	goto L36
L42:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v148 != 0 {
		goto L32
	} else {
		goto L43
	}
L43:
	;
	goto L33
}
func F_logicalrep_worker_wakeup_ptr(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_SetLatch(m, v2+int32(20))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
