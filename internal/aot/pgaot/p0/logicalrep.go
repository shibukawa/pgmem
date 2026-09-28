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
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v11 = v8 + int32(32)
	F_initStringInfo(m, v11)
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
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	m.G0 = v8 + int32(48)
	return v223
L4:
	;
	if v72 < int32(0) {
		goto L3
	} else {
		goto L15
	}
L5:
	;
	v72 = base.I32_ctz(v58) | v59<<(uint(int32(5))%32)
	goto L4
L6:
	;
	v72 = int32(-2)
	goto L4
L7:
	;
	v23 = int32(0)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v26 <= v23 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v29 = l1 + int32(8)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v36 = v33 & int32(-1)
	if v36 != 0 {
		v58 = v36
		v59 = v23
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v37 = int32(1)
	if v37 == v26 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v41 = v37
	goto L11
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v29+v41<<(uint(int32(2))%32))))
	if v48 != 0 {
		v58 = v48
		v59 = v41
		goto L5
	} else {
		goto L13
	}
L12:
	;
	goto L6
L13:
	;
	v50 = v41 + int32(1)
	if v50 != v26 {
		v41 = v50
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+v72<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v79
	F_appendStringInfo(m, v11, int32(_a_F_logicalrep_get_attrs_str_0), v8+int32(16))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if l1 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	if v141 < int32(0) {
		goto L3
	} else {
		goto L28
	}
L18:
	;
	v141 = base.I32_ctz(v127) | v128<<(uint(int32(5))%32)
	goto L17
L19:
	;
	v141 = int32(-2)
	goto L17
L20:
	;
	v92 = v72 + int32(1)
	v94 = int32(base.Ui32(v92) >> (uint(int32(5)) % 32))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v95 <= v94 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v98 = l1 + int32(8)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v98+v94<<(uint(int32(2))%32))))
	v105 = v102 & (int32(-1) << (uint(v92) % 32))
	if v105 != 0 {
		v127 = v105
		v128 = v94
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v107 = v94 + int32(1)
	if v107 == v95 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v110 = v107
	goto L24
L24:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v98+v110<<(uint(int32(2))%32))))
	if v117 != 0 {
		v127 = v117
		v128 = v110
		goto L18
	} else {
		goto L26
	}
L25:
	;
	goto L19
L26:
	;
	v119 = v110 + int32(1)
	if v119 != v95 {
		v110 = v119
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v147 = v141
	goto L29
L29:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v149+v147<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v153
	F_appendStringInfo(m, v8+int32(32), int32(_a_F_logicalrep_get_attrs_str_1), v8)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	goto L3
L31:
	;
	if l1 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	if int32(0) <= v215 {
		v147 = v215
		goto L29
	} else {
		goto L43
	}
L33:
	;
	v215 = base.I32_ctz(v201) | v202<<(uint(int32(5))%32)
	goto L32
L34:
	;
	v215 = int32(-2)
	goto L32
L35:
	;
	v166 = v147 + int32(1)
	v168 = int32(base.Ui32(v166) >> (uint(int32(5)) % 32))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v169 <= v168 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v172 = l1 + int32(8)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v172+v168<<(uint(int32(2))%32))))
	v179 = v176 & (int32(-1) << (uint(v166) % 32))
	if v179 != 0 {
		v201 = v179
		v202 = v168
		goto L33
	} else {
		goto L37
	}
L37:
	;
	v181 = v168 + int32(1)
	if v181 == v169 {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v184 = v181
	goto L39
L39:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v172+v184<<(uint(int32(2))%32))))
	if v191 != 0 {
		v201 = v191
		v202 = v184
		goto L33
	} else {
		goto L41
	}
L40:
	;
	goto L34
L41:
	;
	v193 = v184 + int32(1)
	if v193 != v169 {
		v184 = v193
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	goto L30
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
					v30 = F_dsa_create_ext(m, int32(86), int32(_a_F_logicalrep_launcher_attach_dshmem_2), int32(134217728))
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
				v30 = F_dsa_create_ext(m, int32(86), int32(_a_F_logicalrep_launcher_attach_dshmem_2), int32(134217728))
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	v11 = int32(16)
	v12 = v10 - v11
	m.G0 = v12
	v16 = F_pq_getmsgint(m, l0, int32(2))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v18 = F_palloc0_mul(m, v11, v16)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v18
	v22 = F_palloc_mul(m, int32(1), v16)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v22
	if int32(0) < v16 {
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
	if v85 != v16 {
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
	F_errfinish(m, int32(_a_F_logicalrep_read_tuple_1), int32(915), int32(_a_F_logicalrep_read_tuple_2))
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
func F_logicalrep_relmap_invalidate_cb(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_relmap_invalidate_cb[0]))
	if v11 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(32)
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
	F_hash_seq_init(m, v8+int32(12), v11)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v36 = v8 + int32(12)
	F_hash_seq_init(m, v36, v11)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
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
	v24 = v8 + int32(12)
	v25 = F_hash_seq_search(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v31 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+32)) = uint8(v31)
	F_hash_seq_term(m, v24)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L13
	}
L10:
	;
	if v25 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	if v29 != l1 {
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
	v39 = F_hash_seq_search(m, v36)
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
	v44 = v39
	goto L17
L17:
	;
	v48 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+32)) = uint8(v48)
	v52 = F_hash_seq_search(m, v8+int32(12))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L6
	} else {
		goto L19
	}
L18:
	;
	goto L1
L19:
	;
	if v52 != 0 {
		v44 = v52
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(309237645316)
	v36 = F_hash_create(m, int32(_a_F_logicalrep_relmap_update_3), int64(128), v10, int32(1064))
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
	F_CacheRegisterRelcacheCallback(m, int32(1085))
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
	v76 = F_palloc_mul(m, int32(4), v73)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v76
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v81 = F_palloc_mul(m, int32(4), v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+20)) = v81
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v84 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v91 = int32(0)
	goto L23
L21:
	;
	goto L22
L22:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+24)) = uint8(v121)
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	if v123 != 0 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v95 = v91 << (uint(int32(2)) % 32)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v95+v96)))
	v99 = F_pstrdup(m, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L7
	} else {
		goto L25
	}
L24:
	;
	goto L22
L25:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v101+v95))) = v99
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v106+v95)))
	*(*int32)(unsafe.Add(mBase, uint32(v104+v95))) = v108
	v111 = v91 + int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v111 < v112 {
		v91 = v111
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v125 = v123
	goto L29
L28:
	;
	v125 = int32(114)
	goto L29
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+25)) = uint8(v125)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v128 = F_bms_copy(m, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v128
	*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_relmap_update[3])) = v58
	m.G0 = v10 + int32(48)
	return
}
func F_logicalrep_worker_stop(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop[0]))
	v14 = F_LWLockAcquire(m, v10+int32(_a_F_logicalrep_worker_stop_0), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop[1]))
	if v17 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop[0]))
	F_LWLockRelease(m, v63+int32(_a_F_logicalrep_worker_stop_0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L15
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop[2]))
	v28 = int32(0)
	goto L5
L5:
	;
	v34 = v21 + int32(16) + v28<<(uint(int32(7))%32)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+16)))
	if v35 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	F_logicalrep_worker_stop_internal(m, v34, int32(15))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L14
	}
L7:
	;
	goto L6
L8:
	;
	v49 = v28 + int32(1)
	if v49 != v17 {
		v28 = v49
		goto L5
	} else {
		goto L13
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v38 == int32(4) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)+32))
	if base.B2i32(v41 != l1)|base.B2i32(l0 != v38) != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v34)+36))
	if v45 == l2 {
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
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	v5 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v6|base.B2i32(v7 != int32(1)) != 0 {
		v108 = v6
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	v110 = F_pgmem_kill(m, v109, l1)
	mBase = m.M
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v111 == int32(0) {
		goto L1
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
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[0]))
	v47 = F_LWLockAcquire(m, v43+int32(_a_F_logicalrep_worker_stop_internal_0), int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
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
	v30 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v30
	v35 = base.AtomicRmwOr32(m, v30, int32(_a_F_logicalrep_worker_stop_internal_1), v30)
	goto L9
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[2]))
	if v37 == int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	goto L6
L12:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v49 != int32(1) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v52 != v5 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v54 != 0 {
		v108 = v54
		goto L2
	} else {
		goto L15
	}
L15:
	;
	goto L16
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[0]))
	F_LWLockRelease(m, v60+int32(_a_F_logicalrep_worker_stop_internal_0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	v108 = v102
	goto L2
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[1]))
	v70 = F_WaitLatch(m, v66, int32(41), int32(10), int32(134217734))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L20
	}
L19:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[0]))
	v95 = F_LWLockAcquire(m, v91+int32(_a_F_logicalrep_worker_stop_internal_0), int32(1))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L25
	}
L20:
	;
	if v70&int32(1) == int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[1]))
	v78 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v78
	v83 = base.AtomicRmwOr32(m, v78, int32(_a_F_logicalrep_worker_stop_internal_1), v78)
	goto L22
L22:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[2]))
	if v85 == int32(0) {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	goto L19
L25:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v97 != int32(1) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v100 != v5 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v102 == int32(0) {
		goto L16
	} else {
		goto L28
	}
L28:
	;
	goto L17
L29:
	;
	goto L30
L30:
	;
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v118 != v5 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	goto L1
L32:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[0]))
	F_LWLockRelease(m, v121+int32(_a_F_logicalrep_worker_stop_internal_0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[1]))
	v131 = F_WaitLatch(m, v127, int32(41), int32(10), int32(134217733))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L35
	}
L34:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[0]))
	v156 = F_LWLockAcquire(m, v152+int32(_a_F_logicalrep_worker_stop_internal_0), int32(1))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L40
	}
L35:
	;
	if v131&int32(1) == int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[1]))
	v139 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v139
	v144 = base.AtomicRmwOr32(m, v139, int32(_a_F_logicalrep_worker_stop_internal_1), v139)
	goto L37
L37:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_stop_internal[2]))
	if v146 == int32(0) {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	goto L34
L40:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v158 != 0 {
		goto L30
	} else {
		goto L41
	}
L41:
	;
	goto L31
}
func F_logicalrep_worker_wakeup_ptr(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = v2 + int32(316)
	v5 = int32(0)
	v8 = base.AtomicRmwOr32(m, v5, int32(_a_F_logicalrep_worker_wakeup_ptr_0), v5)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	if v9 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(1)
	v12 = int32(0)
	v15 = base.AtomicRmwOr32(m, v12, int32(_a_F_logicalrep_worker_wakeup_ptr_0), v12)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	if v16 == v12 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	if v19 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_wakeup_ptr[0]))
	if v23 == v19 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v25 = m.G0
	v27 = v25 - int32(16)
	m.G0 = v27
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_wakeup_ptr[1]))
	if v30 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v53 = F_pgmem_kill(m, v19, int32(23))
	mBase = m.M
	goto L2
L9:
	;
	m.G0 = v27 + int32(16)
	goto L1
L10:
	;
	v33 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+15)) = uint8(v33)
	goto L11
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_wakeup_ptr[2]))
	v41 = F_write(m, v37, v27+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v41 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_wakeup_ptr[3]))
	if v45 == int32(27) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
