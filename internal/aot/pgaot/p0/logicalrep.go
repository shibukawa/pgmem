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
	F_appendStringInfoString(m, v8+int32(16), int32(_a_F_logicalrep_get_attrs_str_0))
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
	F_appendStringInfo(m, v8+int32(16), int32(_a_F_logicalrep_get_attrs_str_1), v8)
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	if v5 != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[1]))
		if v7 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[2]))
			v13 = F_LWLockAcquire(m, v9+int32(_a_F_logicalrep_launcher_attach_dshmem_0), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				v15 = int32(_a_F_logicalrep_launcher_attach_dshmem_1)
				v16 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[3]))
				v19 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[4]))
				*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[3])) = v19
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[0]))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
				if v23 == int32(0) {
					v30 = F_dsa_create_ext(m, int32(82), int32(_a_F_logicalrep_launcher_attach_dshmem_2), int32(134217728))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[5])) = v30
						F_dsa_pin(m, v30)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[5]))
							F_dsa_pin_mapping(m, v36)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[5]))
								v44 = F_dshash_create(m, v41, int32(_a_F_logicalrep_launcher_attach_dshmem_3), int32(0))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[1])) = v44
									v48 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[5]))
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+28))
									v52 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[0]))
									*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v50
									v55 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[1]))
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+32))
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
									v59 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[0]))
									*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v57
									*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[3])) = v16
									v85 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[2]))
									F_LWLockRelease(m, v85+int32(_a_F_logicalrep_launcher_attach_dshmem_0))
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
					v62 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[1]))
					if v62 != 0 {
						*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[3])) = v16
						v85 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[2]))
						F_LWLockRelease(m, v85+int32(_a_F_logicalrep_launcher_attach_dshmem_0))
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
							*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[5])) = v65
							F_dsa_pin_mapping(m, v65)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								v72 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[5]))
								v75 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[0]))
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
								v78 = F_dshash_attach(m, v72, int32(_a_F_logicalrep_launcher_attach_dshmem_3), v76, int32(0))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[1])) = v78
									*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[3])) = v16
									v85 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[2]))
									F_LWLockRelease(m, v85+int32(_a_F_logicalrep_launcher_attach_dshmem_0))
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
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[2]))
		v13 = F_LWLockAcquire(m, v9+int32(_a_F_logicalrep_launcher_attach_dshmem_0), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v15 = int32(_a_F_logicalrep_launcher_attach_dshmem_1)
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[3]))
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[4]))
			*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[3])) = v19
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[0]))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
			if v23 == int32(0) {
				v30 = F_dsa_create_ext(m, int32(82), int32(_a_F_logicalrep_launcher_attach_dshmem_2), int32(134217728))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[5])) = v30
					F_dsa_pin(m, v30)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[5]))
						F_dsa_pin_mapping(m, v36)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[5]))
							v44 = F_dshash_create(m, v41, int32(_a_F_logicalrep_launcher_attach_dshmem_3), int32(0))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[1])) = v44
								v48 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[5]))
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+28))
								v52 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[0]))
								*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v50
								v55 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[1]))
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+32))
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
								v59 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[0]))
								*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v57
								*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[3])) = v16
								v85 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[2]))
								F_LWLockRelease(m, v85+int32(_a_F_logicalrep_launcher_attach_dshmem_0))
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
				v62 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[1]))
				if v62 != 0 {
					*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[3])) = v16
					v85 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[2]))
					F_LWLockRelease(m, v85+int32(_a_F_logicalrep_launcher_attach_dshmem_0))
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
						*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[5])) = v65
						F_dsa_pin_mapping(m, v65)
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							v72 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[5]))
							v75 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[0]))
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
							v78 = F_dshash_attach(m, v72, int32(_a_F_logicalrep_launcher_attach_dshmem_3), v76, int32(0))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[1])) = v78
								*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[3])) = v16
								v85 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_attach_dshmem[2]))
								F_LWLockRelease(m, v85+int32(_a_F_logicalrep_launcher_attach_dshmem_0))
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
	F_errmsg_internal(m, int32(_a_F_logicalrep_read_tuple_0), v12)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_logicalrep_read_tuple_1), int32(912), int32(_a_F_logicalrep_read_tuple_2))
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
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_relmap_invalidate_cb[0]))
	if v10 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v7 + int32(32)
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
	F_hash_seq_init(m, v7+int32(12), v10)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v34 = v7 + int32(12)
	F_hash_seq_init(m, v34, v10)
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
	v22 = v7 + int32(12)
	v23 = F_hash_seq_search(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v29 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+32)) = uint8(v29)
	F_hash_seq_term(m, v22)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L13
	}
L10:
	;
	if v23 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+36))
	if v27 != l1 {
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
	v37 = F_hash_seq_search(m, v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	if v37 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v41 = v37
	goto L17
L17:
	;
	v45 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+32)) = uint8(v45)
	v49 = F_hash_seq_search(m, v7+int32(12))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L19
	}
L18:
	;
	goto L1
L19:
	;
	if v49 != 0 {
		v41 = v49
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
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_relmap_update[0]))
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
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_relmap_update[1]))
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
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_relmap_update[2]))
	v25 = F_AllocSetContextCreateInternal(m, v20, int32(_a_F_logicalrep_relmap_update_0), int32(0), int32(_a_F_logicalrep_relmap_update_1), int32(_a_F_logicalrep_relmap_update_2))
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
	v36 = F_hash_create(m, int32(_a_F_logicalrep_relmap_update_3), int32(128), v10, int32(1064))
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
	*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_relmap_update[1])) = v25
	v28 = v25
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_relmap_update[0])) = v36
	F_CacheRegisterRelcacheCallback(m, int32(1015))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_relmap_update[0]))
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
	base.MemoryFill(m, v47, int32(0), int32(72))
	v57 = int32(_a_F_logicalrep_relmap_update_4)
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_relmap_update[3]))
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_relmap_update[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_relmap_update[3])) = v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v66 = F_pstrdup(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L7
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v66
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v70 = F_pstrdup(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v73
	v77 = F_palloc(m, v73<<(uint(int32(2))%32))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v77
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v83 = F_palloc(m, v80<<(uint(int32(2))%32))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+20)) = v83
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v86 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v93 = int32(0)
	goto L23
L21:
	;
	goto L22
L22:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+24)) = uint8(v123)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v126 = F_bms_copy(m, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L7
	} else {
		goto L27
	}
L23:
	;
	v97 = v93 << (uint(int32(2)) % 32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97+v98)))
	v101 = F_pstrdup(m, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L25
	}
L24:
	;
	goto L22
L25:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v103+v97))) = v101
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v108+v97)))
	*(*int32)(unsafe.Add(mBase, uint32(v106+v97))) = v110
	v113 = v93 + int32(1)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v113 < v114 {
		v93 = v113
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v126
	*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_relmap_update[3])) = v58
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
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop[0]))
	v12 = F_LWLockAcquire(m, v8+int32(_a_F_logicalrep_worker_stop_0), int32(1))
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
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop[1]))
	if v15 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop[0]))
	F_LWLockRelease(m, v54+int32(_a_F_logicalrep_worker_stop_0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L15
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop[2]))
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
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	v5 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v6|base.B2i32(v7 != int32(1)) != 0 {
		v100 = v6
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+44))
	v102 = F_kill(m, v101, l1)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L29
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[0]))
	F_LWLockRelease(m, v12+int32(_a_F_logicalrep_worker_stop_internal_0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[1]))
	v22 = F_WaitLatch(m, v18, int32(41), int32(10), int32(134217734))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L7
	}
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[0]))
	v43 = F_LWLockAcquire(m, v39+int32(_a_F_logicalrep_worker_stop_internal_0), int32(1))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L12
	}
L7:
	;
	if v22&int32(1) == int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(0)
	goto L9
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[2]))
	if v33 == int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	goto L6
L12:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v45 != int32(1) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v48 != v5 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v50 != 0 {
		v100 = v50
		goto L2
	} else {
		goto L15
	}
L15:
	;
	goto L16
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[0]))
	F_LWLockRelease(m, v56+int32(_a_F_logicalrep_worker_stop_internal_0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	v100 = v94
	goto L2
L18:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[1]))
	v66 = F_WaitLatch(m, v62, int32(41), int32(10), int32(134217734))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L20
	}
L19:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[0]))
	v87 = F_LWLockAcquire(m, v83+int32(_a_F_logicalrep_worker_stop_internal_0), int32(1))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L25
	}
L20:
	;
	if v66&int32(1) == int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(0)
	goto L22
L22:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[2]))
	if v77 == int32(0) {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	goto L19
L25:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v89 != int32(1) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v92 != v5 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v94 == int32(0) {
		goto L16
	} else {
		goto L28
	}
L28:
	;
	goto L17
L29:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v104 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L31
L31:
	;
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v111 != v5 {
		goto L1
	} else {
		goto L33
	}
L32:
	;
	goto L1
L33:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[0]))
	F_LWLockRelease(m, v114+int32(_a_F_logicalrep_worker_stop_internal_0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[1]))
	v124 = F_WaitLatch(m, v120, int32(41), int32(10), int32(134217733))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L36
	}
L35:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[0]))
	v145 = F_LWLockAcquire(m, v141+int32(_a_F_logicalrep_worker_stop_internal_0), int32(1))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L41
	}
L36:
	;
	if v124&int32(1) == int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = int32(0)
	goto L38
L38:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[2]))
	if v135 == int32(0) {
		goto L35
	} else {
		goto L39
	}
L39:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	goto L35
L41:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v147 != 0 {
		goto L31
	} else {
		goto L42
	}
L42:
	;
	goto L32
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
