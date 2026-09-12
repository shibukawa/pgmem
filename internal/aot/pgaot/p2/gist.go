package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gistPopItupFromNodeBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v13 <= int32(0) {
		m.G0 = v11 + int32(16)
		return base.B2i32(int32(0) < v13)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		if v16 == int32(0) {
			F_gistLoadNodeBuffer(m, l0, l1)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v24 = v23
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
				v26 = v24 + v25
				v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+14)))
				v29 = v27 & int32(8191)
				v30 = F_palloc(m, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v30
					if v29 != 0 {
						v35 = F__emscripten_memcpy_bulkmem(m, v30, v26+int32(8), v29)
						mBase = m.M
					} else {
					}
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v37 + (v29+int32(7))&int32(16376)
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
					if v45 != int32(8184) {
						m.G0 = v11 + int32(16)
						return base.B2i32(int32(0) < v13)
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v48 - int32(1)
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
						if v52 != int32(-1) {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v57 = F_BufFileSeekBlock(m, v55, base.I64_extend_i32_s(v52))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								if v57 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = v52
										F_errmsg_internal(m, int32(405034), v11)
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(516218), int32(753), int32(331981))
											mBase = m.M
											v112 = m.ExcPending
											if v112 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									F_BufFileReadExact(m, v55, v44, int32(8192))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										if v62 < v63 {
											v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v76 = v62
											v77 = v65
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v76 + int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v77+v76<<(uint(int32(2))%32)))) = v52
											m.G0 = v11 + int32(16)
											return base.B2i32(int32(0) < v13)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v63 << (uint(int32(1)) % 32)
											v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v72 = F_repalloc(m, v69, v63<<(uint(int32(3))%32))
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v72
												v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												v76 = v75
												v77 = v72
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v76 + int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v77+v76<<(uint(int32(2))%32)))) = v52
												m.G0 = v11 + int32(16)
												return base.B2i32(int32(0) < v13)
											}
										}
									}
								}
							}
						} else {
							F_pfree(m, v44)
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(0)
								m.G0 = v11 + int32(16)
								return base.B2i32(int32(0) < v13)
							}
						}
					}
				}
			}
		} else {
			v24 = v16
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
			v26 = v24 + v25
			v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+14)))
			v29 = v27 & int32(8191)
			v30 = F_palloc(m, v29)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v30
				if v29 != 0 {
					v35 = F__emscripten_memcpy_bulkmem(m, v30, v26+int32(8), v29)
					mBase = m.M
				} else {
				}
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v37 + (v29+int32(7))&int32(16376)
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
				if v45 != int32(8184) {
					m.G0 = v11 + int32(16)
					return base.B2i32(int32(0) < v13)
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v48 - int32(1)
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
					if v52 != int32(-1) {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v57 = F_BufFileSeekBlock(m, v55, base.I64_extend_i32_s(v52))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							if v57 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v52
									F_errmsg_internal(m, int32(405034), v11)
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(516218), int32(753), int32(331981))
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								F_BufFileReadExact(m, v55, v44, int32(8192))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									if v62 < v63 {
										v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v76 = v62
										v77 = v65
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v76 + int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v77+v76<<(uint(int32(2))%32)))) = v52
										m.G0 = v11 + int32(16)
										return base.B2i32(int32(0) < v13)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v63 << (uint(int32(1)) % 32)
										v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v72 = F_repalloc(m, v69, v63<<(uint(int32(3))%32))
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v72
											v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											v76 = v75
											v77 = v72
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v76 + int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v77+v76<<(uint(int32(2))%32)))) = v52
											m.G0 = v11 + int32(16)
											return base.B2i32(int32(0) < v13)
										}
									}
								}
							}
						}
					} else {
						F_pfree(m, v44)
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(0)
							m.G0 = v11 + int32(16)
							return base.B2i32(int32(0) < v13)
						}
					}
				}
			}
		}
	}
}
func F_gistProcessEmptyingQueue(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = v16
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v13 + int32(16)
	return
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v29 = F_list_delete_first(m, v18)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v29
	v32 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+16)) = uint8(v32)
	v35 = m.G0
	v37 = v35 - int32(16)
	m.G0 = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	if v32 < v39 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v125 = F_gistPopItupFromNodeBuffer(m, v15, v28, v13+int32(12))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L6
	} else {
		goto L31
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L6
	} else {
		goto L27
	}
L10:
	;
	v43 = v39
	v48 = v32
	goto L13
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = int32(0)
	m.G0 = v37 + int32(16)
	goto L8
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52+v48<<(uint(int32(2))%32))))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v57 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L12
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if int32(0) < v58 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v89 = v43
	goto L17
L17:
	;
	v93 = v48 + int32(1)
	if v93 < v89 {
		v43 = v89
		v48 = v93
		goto L13
	} else {
		goto L26
	}
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v77 = F_BufFileSeekBlock(m, v75, base.I64_extend_i32_s(v73))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L6
	} else {
		goto L22
	}
L19:
	;
	v62 = v58 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v62<<(uint(int32(2))%32))))
	v73 = v68
	goto L18
L20:
	;
	goto L21
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v69 + int32(1)
	v73 = v69
	goto L18
L22:
	;
	if v77 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	F_BufFileWrite(m, v75, v74, int32(8192))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L6
	} else {
		goto L24
	}
L24:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	F_pfree(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v56)+12)) = int32(0)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v89 = v88
	goto L17
L26:
	;
	goto L14
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v73
	F_errmsg_internal(m, int32(405034), v37)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(516218), int32(761), int32(331962))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	if v162 != 0 {
		v18 = v162
		goto L4
	} else {
		goto L40
	}
L31:
	;
	if v125 == int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	goto L33
L33:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	v142 = F_gistProcessItup(m, l0, v139, v140, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L6
	} else {
		goto L35
	}
L34:
	;
	goto L30
L35:
	;
	if v142 != 0 {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	F_MemoryContextReset(m, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v150 = F_gistPopItupFromNodeBuffer(m, v15, v28, v13+int32(12))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	if v150 != 0 {
		goto L33
	} else {
		goto L39
	}
L39:
	;
	goto L34
L40:
	;
	goto L5
}
func F_gistPushItupToNodeBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
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
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(4554240)
	v14 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v54 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	v20 = F_MemoryContextAllocZero(m, v16, int32(8192))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = int64(35154307317759)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v20
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
	if v27 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v28 < v29 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43+v42<<(uint(int32(2))%32)))) = l1
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v48 + int32(1)
	goto L1
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v42 = v28
	v43 = v31
	goto L6
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v29 << (uint(int32(1)) % 32)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v38 = F_repalloc(m, v35, v29<<(uint(int32(3))%32))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v38
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v42 = v41
	v43 = v38
	goto L6
L11:
	;
	F_gistLoadNodeBuffer(m, l0, l1)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L3
	} else {
		goto L14
	}
L12:
	;
	v60 = v54
	goto L13
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
	v64 = v62 & int32(8191)
	if base.Ui32(v61) < base.Ui32((v64+int32(7))&int32(16376)) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v60 = v59
	goto L13
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L3
	} else {
		goto L34
	}
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(0) < v70 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v108 = v60
	v109 = v64
	v110 = v61
	goto L18
L18:
	;
	v115 = v110 - (v109+int32(7))&int32(16376)
	*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v115
	if v109 != 0 {
		goto L27
	} else {
		goto L28
	}
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v89 = F_BufFileSeekBlock(m, v87, base.I64_extend_i32_s(v85))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L3
	} else {
		goto L23
	}
L20:
	;
	v74 = v70 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76+v74<<(uint(int32(2))%32))))
	v85 = v80
	goto L19
L21:
	;
	goto L22
L22:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v81 + int32(1)
	v85 = v81
	goto L19
L23:
	;
	if v89 != 0 {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	F_BufFileWrite(m, v87, v86, int32(8192))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = int32(8184)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v85
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v99 + int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
	v108 = v103
	v109 = v105 & int32(8191)
	v110 = v104
	goto L18
L26:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v125 = base.I32_div_s(v123, int32(2))
	if v122 <= v125 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v120 = F__emscripten_memcpy_bulkmem(m, v108+v115+int32(8), l2, v109)
	mBase = m.M
	goto L29
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v14
	m.G0 = v11 + int32(16)
	return
L31:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v127 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v129 = F_lcons(m, l1, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v129
	v132 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v132)
	goto L30
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v85
	F_errmsg_internal(m, int32(405034), v11)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(516218), int32(761), int32(331962))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gist_bbox_zorder_abbrev_abort(m *base.Module, l0 int32, l1 int32) int32 {
	return int32(0)
}
func F_gist_page_items(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v462 int32
	_ = v462
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	v17 = m.G0
	v19 = v17 - int32(256)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v28 = F_superuser(m)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L100
	}
L4:
	;
	if v28 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L96
	}
L8:
	;
	v34 = F_index_open(m, v27, int32(1))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+84))
	if v37 != int32(783) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v40 = F_verify_gist_page(m, v22)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	m.G0 = v19 + int32(256)
	return int32(0)
L12:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+14)))
	if v42 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_relation_close(m, v34, int32(1))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v50 = int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+16)))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v52)+12)))
	if v54&v50 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v48 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v48)
	goto L11
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v34)+192))
	v61 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+10)))
	v62 = F_CreateTupleDescTruncatedCopy(m, v51, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	v64 = v50
	v65 = v51
	goto L19
L19:
	;
	v66 = int32(0)
	v68 = int32(1)
	v69 = int32(2)
	if v64&v68 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v64 = int32(3)
	v65 = v62
	goto L19
L21:
	;
	v79 = int32(7)
	goto L23
L22:
	;
	v79 = v69
	goto L23
L23:
	;
	v81 = F_pg_get_indexdef_worker(m, v27, v66, v66, v68, int32(base.Ui32(v64&v69)>>(uint(v68)%32)), v66, v66, v79, int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+16)))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v83)+12)))
	if v85&int32(2) != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	F_relation_close(m, v34, int32(1))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L95
	}
L26:
	;
	v90 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+12)))
	if base.Ui32(v103) < base.Ui32(int32(25)) {
		goto L25
	} else {
		goto L33
	}
L29:
	;
	if v90 == int32(0) {
		goto L25
	} else {
		goto L30
	}
L30:
	;
	F_errmsg_internal(m, int32(467150), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(517468), int32(255), int32(159926))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L25
L33:
	;
	v107 = v103 + int32(262120)
	if v107&int32(262140) == int32(0) {
		goto L25
	} else {
		goto L34
	}
L34:
	;
	v122 = int32(1)
	v124 = v122
	v135 = v122
	goto L35
L35:
	;
	v144 = v124<<(uint(int32(2))%32) + (v40 + int32(24)) - int32(4)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	v148 = v40 + v145&int32(32767)
	F_index_deform_tuple(m, v148, v65, v19+int32(80), v19+int32(48))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L37
	}
L36:
	;
	goto L25
L37:
	;
	v155 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(220)))) = uint8(v155)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+216)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v19)+228)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v19)+224)) = v124
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v148)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+232)) = v161 & int32(8191)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	v166 = int32(98304)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+236)) = base.B2i32(v165&v166 == v166)
	if v81 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_initStringInfo(m, v19+int32(32))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	v413 = v155
	v415 = int32(1)
	goto L40
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+220)) = uint8(v415)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+240)) = v413
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	F_tuplestore_putvalues(m, v431, v432, v19+int32(224), v19+int32(216))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L93
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v81
	F_appendStringInfo(m, v19+int32(32), int32(715898), v19)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v183 = int32(0)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v183 < v184 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v187 = v184
	v193 = v183
	goto L46
L44:
	;
	goto L45
L45:
	;
	F_appendStringInfoChar(m, v19+int32(32), int32(41))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L91
	}
L46:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(48)+v193))))
	if v206 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L45
L48:
	;
	v230 = int32(317374)
	goto L50
L49:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v65+int32(88)+v187<<(uint(int32(4))%32)+v193*int32(100))))
	F_getTypeOutputInfo(m, v214, v19+int32(28), v19+int32(27))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L51
	}
L50:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v34)+192))
	v234 = int32(*(*int16)(unsafe.Add(mBase, uint32(v233)+10)))
	if v234 == v193 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(80)+v193<<(uint(int32(2))%32))))
	v228 = F_OidOutputFunctionCall(m, v221, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v230 = v228
	goto L50
L53:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	v246 = v230
	v247 = v243
	goto L61
L54:
	;
	v240 = int32(716856)
	goto L56
L55:
	;
	if v193 == int32(0) {
		goto L53
	} else {
		goto L57
	}
L56:
	;
	F_appendStringInfoString(m, v19+int32(32), v240)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L58
	}
L57:
	;
	v240 = int32(778962)
	goto L56
L58:
	;
	goto L53
L59:
	;
	v292 = v230
	goto L71
L60:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	if v266 <= v267+int32(1) {
		goto L66
	} else {
		goto L67
	}
L61:
	;
	switch v247 & int32(255) {
	case 0:
		goto L63
	default:
		goto L64
	case 9, 10, 11, 12, 13, 32, 34, 40, 41, 44, 92:
		goto L60
	}
L62:
	;
	if v243 != 0 {
		v291 = int32(0)
		goto L59
	} else {
		goto L65
	}
L63:
	;
	goto L62
L64:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+1)))
	v246 = v246 + int32(1)
	v247 = v262
	goto L61
L65:
	;
	goto L60
L66:
	;
	F_appendStringInfoChar(m, v19+int32(32), int32(34))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v279 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v277+v267))) = uint8(v279)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	v282 = int32(1)
	v283 = v281 + v282
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v283
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v287 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v285+v283))) = uint8(v287)
	v291 = v282
	goto L59
L69:
	;
	v291 = int32(1)
	goto L59
L70:
	;
	v385 = v193 + int32(1)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v385 < v386 {
		v187 = v386
		v193 = v385
		goto L46
	} else {
		goto L90
	}
L71:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
	v309 = base.I32_extend8_s(v308)
	if v308 == int32(34) {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v373 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v371+v317))) = uint8(v373)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	v377 = v375 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v377
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v381 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v379+v377))) = uint8(v381)
	goto L70
L73:
	;
	goto L72
L74:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	if v347 <= v348+int32(1) {
		goto L86
	} else {
		goto L87
	}
L75:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	if v326 <= v327+int32(1) {
		goto L82
	} else {
		goto L83
	}
L76:
	;
	if v308 == int32(92) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	if v308 != 0 {
		goto L74
	} else {
		goto L78
	}
L78:
	;
	if v291 == int32(0) {
		goto L70
	} else {
		goto L79
	}
L79:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	if v317+int32(1) < v316 {
		goto L73
	} else {
		goto L80
	}
L80:
	;
	F_appendStringInfoChar(m, v19+int32(32), int32(34))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	goto L70
L82:
	;
	F_appendStringInfoChar(m, v19+int32(32), v309)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v335+v327))) = uint8(v309)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	v340 = v338 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v340
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v344 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v342+v340))) = uint8(v344)
	goto L74
L85:
	;
	goto L74
L86:
	;
	F_appendStringInfoChar(m, v19+int32(32), v309)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v358+v348))) = uint8(v309)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	v362 = int32(1)
	v363 = v361 + v362
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v363
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v367 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v365+v363))) = uint8(v367)
	v292 = v292 + v362
	goto L71
L89:
	;
	v292 = v292 + int32(1)
	goto L71
L90:
	;
	goto L47
L91:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v411 = F_cstring_to_text(m, v410)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v413 = v411
	v415 = int32(0)
	goto L40
L93:
	;
	v440 = v135 + int32(1)
	v442 = v440 & int32(65535)
	if base.Ui32(v442) <= base.Ui32(int32(base.Ui32(v107)>>(uint(int32(2))%32))&int32(65535)) {
		v124 = v442
		v135 = v440
		goto L35
	} else {
		goto L94
	}
L94:
	;
	goto L36
L95:
	;
	goto L11
L96:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errmsg(m, int32(149499), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(517468), int32(211), int32(159926))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = int32(541708)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v507 + int32(4)
	F_errmsg(m, int32(29213), v19+int32(16))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(517468), int32(222), int32(159926))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gist_page_items_bytea(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v22 = F_superuser(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L31
	}
L7:
	;
	v27 = F_verify_gist_page(m, v17)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	m.G0 = v14 + int32(48)
	return int32(0)
L9:
	;
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+14)))
	if v29 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v32 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
	goto L8
L11:
	;
	goto L12
L12:
	;
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+16)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v34)+12)))
	if v36&int32(2) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v41 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+12)))
	if base.Ui32(v58) < base.Ui32(int32(25)) {
		goto L8
	} else {
		goto L20
	}
L16:
	;
	if v41 == int32(0) {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	F_errmsg_internal(m, int32(467150), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(517468), int32(152), int32(531426))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	goto L8
L20:
	;
	v62 = v58 + int32(262120)
	if v62&int32(262140) == int32(0) {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v73 = int32(1)
	v75 = v73
	v81 = v73
	goto L22
L22:
	;
	v89 = int32(4)
	v90 = v75<<(uint(int32(2))%32) + (v27 + int32(24)) - v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v94 = v27 + v91&int32(32767)
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+6)))
	v96 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+12)) = uint8(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v75
	v103 = v95 & int32(8191)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v103
	v106 = v103 + v89
	v107 = F_palloc(m, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	goto L8
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v106 << (uint(int32(2)) % 32)
	if v103 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v107
	v118 = int32(98304)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = base.B2i32(v116&v118 == v118)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	F_tuplestore_putvalues(m, v123, v124, v14+int32(16), v14+int32(8))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L29
	}
L26:
	;
	v114 = F__emscripten_memcpy_bulkmem(m, v107+int32(4), v94, v103)
	mBase = m.M
	goto L28
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	v132 = v81 + int32(1)
	v134 = v132 & int32(65535)
	if base.Ui32(v134) <= base.Ui32(int32(base.Ui32(v62)>>(uint(int32(2))%32))&int32(65535)) {
		v75 = v134
		v81 = v132
		goto L22
	} else {
		goto L30
	}
L30:
	;
	goto L23
L31:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errmsg(m, int32(149499), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(517468), int32(141), int32(531426))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gist_page_opaque_info(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v120 int64
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int64
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = F_superuser(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L49
	}
L4:
	;
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v17 = F_verify_gist_page(m, v11)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L45
	}
L8:
	;
	m.G0 = v8 + int32(112)
	return v148
L9:
	;
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+14)))
	if v19 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v22 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
	v148 = int32(0)
	goto L8
L11:
	;
	goto L12
L12:
	;
	v25 = int32(0)
	v29 = F_get_call_result_type(m, l0, v25, v8+int32(108))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v29 != int32(1) {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)))
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+v33)+12)))
	if v35&int32(1) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v41 = F_cstring_to_text(m, int32(356119))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v45 = v25
	v46 = v8
	goto L17
L17:
	;
	if v35&int32(2) != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v41
	v45 = int32(1)
	v46 = v8 | int32(4)
	goto L17
L19:
	;
	v50 = F_cstring_to_text(m, int32(467190))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	v55 = v45
	goto L21
L21:
	;
	if v35&int32(4) != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v50
	v55 = v45 + int32(1)
	goto L21
L23:
	;
	v62 = F_cstring_to_text(m, int32(467037))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	v67 = v55
	goto L25
L25:
	;
	if v35&int32(8) != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8|v55<<(uint(int32(2))%32)))) = v62
	v67 = v55 + int32(1)
	goto L25
L27:
	;
	v74 = F_cstring_to_text(m, int32(111253))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	v79 = v67
	goto L29
L29:
	;
	if v35&int32(16) != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8+v67<<(uint(int32(2))%32)))) = v74
	v79 = v67 + int32(1)
	goto L29
L31:
	;
	v86 = F_cstring_to_text(m, int32(427504))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	v91 = v79
	goto L33
L33:
	;
	v93 = v35 & int32(65504)
	if v93 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8+v79<<(uint(int32(2))%32)))) = v86
	v91 = v79 + int32(1)
	goto L33
L35:
	;
	v99 = F_DirectFunctionCall1Coll(m, int32(2911), int32(0), v93)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	v104 = v91
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+76)) = int32(0)
	v107 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v17)+4)))
	v108 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v17))))
	v112 = F_Int64GetDatum(m, v107|v108<<(uint(int64(32))%64))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8+v91<<(uint(int32(2))%32)))) = v99
	v104 = v91 + int32(1)
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v112
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)))
	v116 = v17 + v115
	v117 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v116))))
	v120 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v116)+4)))
	v122 = F_Int64GetDatum(m, v117<<(uint(int64(32))%64)|v120)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+84)) = v122
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)))
	v127 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v17+v125)+8)))
	v128 = F_Int64GetDatum(m, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+88)) = v128
	v132 = F_construct_array_builtin(m, v8, v104, int32(25))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+92)) = v132
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v8)+108))
	v140 = F_heap_form_tuple(m, v135, v8+int32(80), v8+int32(76))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
	v143 = F_HeapTupleHeaderGetDatum(m, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v148 = v143
	goto L8
L45:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(149499), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(517468), int32(86), int32(253506))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_errmsg_internal(m, int32(384879), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(517468), int32(95), int32(253506))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gist_point_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+20)))
	if v5 == int32(1) {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+32)) = int32(113)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+28)) = int32(114)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(115)
		v15 = int32(116)
	} else {
		v15 = int32(113)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = v15
	return int32(0)
}
func F_gist_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
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
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int64
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int64
	_ = v273
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v326 int64
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int64
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int64
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int64
	_ = v356
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v482 int32
	_ = v482
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v652 int64
	_ = v652
	var v654 int32
	_ = v654
	var v656 int64
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v672 int32
	_ = v672
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v688 int64
	_ = v688
	var v690 int32
	_ = v690
	var v692 int64
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v741 int32
	_ = v741
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int64
	_ = v765
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int64
	_ = v773
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v804 int64
	_ = v804
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v892 int32
	_ = v892
	var v898 int32
	_ = v898
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v913 int32
	_ = v913
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v931 int32
	_ = v931
	v2 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(96)
	m.G0 = v27
	v29 = int32(4554240)
	v30 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+48)))
	v35 = *(*int32)(unsafe.Add(mBase, _consts[54]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v35
	v38 = v32 & int32(240)
	switch int32(base.Ui32(v38) >> (uint(int32(4)) % 32)) {
	case 0:
		goto L9
	case 1:
		goto L8
	case 2:
		goto L7
	case 3:
		goto L6
	default:
		goto L1
	case 6:
		goto L5
	case 7:
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L10
	} else {
		goto L173
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L10
	} else {
		goto L170
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L10
	} else {
		goto L167
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v30
	v883 = *(*int32)(unsafe.Add(mBase, _consts[54]))
	F_MemoryContextReset(m, v883)
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L10
	} else {
		goto L166
	}
L5:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v31)+64))
	v765 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v769 = F_XLogReadBufferForRedo(m, l0, int32(0), v27+int32(92))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L10
	} else {
		goto L141
	}
L6:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v31)+64))
	v355 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v354)+18)))
	if v355 != 0 {
		goto L75
	} else {
		goto L76
	}
L7:
	;
	v340 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	if base.Ui32(v340) < base.Ui32(int32(2)) {
		goto L4
	} else {
		goto L73
	}
L8:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v31)+64))
	v257 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v259 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	if base.Ui32(int32(2)) <= base.Ui32(v259) {
		goto L56
	} else {
		goto L57
	}
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v31)+64))
	v42 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v46 = F_XLogReadBufferForRedo(m, l0, int32(0), v27+int32(76))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	if v46 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v50 = int32(0)
	v52 = v27 + int32(92)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+72))
	if v55 < v50 {
		v77 = v50
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)+72))
	if v242 <= int32(0) {
		goto L50
	} else {
		goto L51
	}
L15:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	if v81 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L16:
	;
	v80 = v77
	goto L15
L17:
	;
	v61 = v54 + int32(76)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v62 != int32(1) {
		v77 = v50
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+43)))
	if v65 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v52 == int32(0) {
		v77 = v50
		goto L16
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v52 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v70 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v70
	v80 = v70
	goto L15
L23:
	;
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v73
	goto L25
L24:
	;
	goto L25
L25:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v61)+44))
	v77 = v75
	goto L16
L26:
	;
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41))))
	switch v100 {
	case 0:
		v130 = v80
		goto L30
	case 1:
		goto L32
	default:
		goto L31
	}
L27:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v85+(v81^int32(-1))<<(uint(int32(2))%32))))
	v99 = v91
	goto L26
L28:
	;
	goto L29
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v99 = v93 + v81<<(uint(int32(13))%32) + int32(-8192)
	goto L26
L30:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
	if base.Ui32(v130-v80) < base.Ui32(v133) {
		goto L38
	} else {
		goto L39
	}
L31:
	;
	F_PageIndexMultiDelete(m, v99, v80, v100)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L10
	} else {
		goto L36
	}
L32:
	;
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+2)))
	if v101 != int32(1) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80))))
	v106 = v80 + int32(2)
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+8)))
	v109 = v107 & int32(8191)
	v110 = F_PageIndexTupleOverwrite(m, v99, v104, v106, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	if v110 == int32(0) {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v130 = v106 + v109
	goto L30
L36:
	;
	v117 = int32(1)
	v119 = v80 + v100<<(uint(v117)%32)
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+16)))
	v121 = v99 + v120
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121)+12)))
	if v122&v117 == int32(0) {
		v130 = v119
		goto L30
	} else {
		goto L37
	}
L37:
	;
	v128 = v122 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v121)+12)) = uint16(v128)
	v130 = v119
	goto L30
L38:
	;
	v136 = int32(1)
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+12)))
	if base.Ui32(v137) < base.Ui32(int32(25)) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v99))) = base.I64_rotr(v42, int64(32))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	F_MarkBufferDirty(m, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L10
	} else {
		goto L49
	}
L41:
	;
	v146 = v136
	goto L43
L42:
	;
	v146 = int32(base.Ui32(v137+int32(262120))>>(uint(int32(2))%32)) + v136
	goto L43
L43:
	;
	v148 = v130
	v150 = v146
	goto L44
L44:
	;
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v148)+6)))
	v173 = v171 & int32(8191)
	v177 = F_PageAddItemExtended(m, v99, v148, v173, v150&int32(65535), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L10
	} else {
		goto L46
	}
L45:
	;
	goto L40
L46:
	;
	if v177 == int32(0) {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
	v184 = v148 + v173
	if base.Ui32(v184-v80) < base.Ui32(v183) {
		v148 = v184
		v150 = v150 + int32(1)
		goto L44
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	goto L14
L50:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	if v251 == int32(0) {
		goto L4
	} else {
		goto L54
	}
L51:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+128)))
	if v245 != int32(1) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	F_gistRedoClearFollowRight(m, l0, int32(1))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L10
	} else {
		goto L53
	}
L53:
	;
	goto L50
L54:
	;
	F_UnlockReleaseBuffer(m, v251)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	goto L4
L56:
	;
	v262 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v262, v27+int32(76), v262, v262)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L10
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v284 = F_XLogReadBufferForRedo(m, l0, int32(0), v27+int32(76))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L10
	} else {
		goto L61
	}
L59:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+6)))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v27)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+56)) = v271
	v273 = *(*int64)(unsafe.Add(mBase, uint32(v27)+76))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+48)) = v273
	F_ResolveRecoveryConflictWithSnapshot(m, v270, v269, v27+int32(48))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	if v284 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	if v290 < int32(0) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	goto L64
L64:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	if v334 == int32(0) {
		goto L4
	} else {
		goto L71
	}
L65:
	;
	v309 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256)+4)))
	F_PageIndexMultiDelete(m, v308, v256+int32(8), v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L10
	} else {
		goto L69
	}
L66:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v294+(v290^int32(-1))<<(uint(int32(2))%32))))
	v308 = v300
	goto L65
L67:
	;
	goto L68
L68:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v308 = v302 + v290<<(uint(int32(13))%32) + int32(-8192)
	goto L65
L69:
	;
	v312 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v308)+16)))
	v313 = v308 + v312
	v314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v313)+12)))
	v316 = v314 & int32(65519)
	*(*uint16)(unsafe.Add(mBase, uint32(v313)+12)) = uint16(v316)
	v318 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v308)+16)))
	v319 = v308 + v318
	v320 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v319)+12)))
	v322 = v320 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v319)+12)) = uint16(v322)
	*(*uint32)(unsafe.Add(mBase, uint32(v308)+4)) = uint32(v257)
	v326 = int64(base.Ui64(v257) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v308))) = uint32(v326)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	F_MarkBufferDirty(m, v328)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L10
	} else {
		goto L70
	}
L70:
	;
	goto L64
L71:
	;
	F_UnlockReleaseBuffer(m, v334)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L10
	} else {
		goto L72
	}
L72:
	;
	goto L4
L73:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v31)+64))
	v344 = *(*int64)(unsafe.Add(mBase, uint32(v343)+16))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+24)))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v343)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+72)) = v346
	v348 = *(*int64)(unsafe.Add(mBase, uint32(v343)))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+64)) = v348
	F_ResolveRecoveryConflictWithSnapshotFullXid(m, v344, v345, v27-int32(-64))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	goto L4
L75:
	;
	v356 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v370 = v2
	v373 = v2
	v374 = v2
	goto L78
L76:
	;
	v730 = v31
	v741 = v2
	goto L77
L77:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v730)+72))
	if v753 < int32(0) {
		goto L136
	} else {
		goto L137
	}
L78:
	;
	v386 = v370 + int32(1)
	v388 = v386 & int32(255)
	v389 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v388, v389, v389, v27+int32(92))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L10
	} else {
		goto L80
	}
L79:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v730 = v728
	v741 = v725
	goto L77
L80:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
	v396 = F_XLogInitBufferForRedo(m, l0, v388)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L10
	} else {
		goto L82
	}
L81:
	;
	v416 = int32(0)
	v419 = v27 + int32(76)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v421)+72))
	if v422 < v388 {
		v444 = v416
		goto L87
	} else {
		goto L88
	}
L82:
	;
	if v396 < int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v401 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v401+(v396^int32(-1))<<(uint(int32(2))%32))))
	v415 = v407
	goto L81
L84:
	;
	goto L85
L85:
	;
	v409 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v415 = v409 + v396<<(uint(int32(13))%32) + int32(-8192)
	goto L81
L86:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	v451 = F_palloc(m, v448<<(uint(int32(2))%32))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L10
	} else {
		goto L97
	}
L87:
	;
	v447 = v444
	goto L86
L88:
	;
	v428 = v421 + v388*int32(52) + int32(76)
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428))))
	if v429 != int32(1) {
		v444 = v416
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428)+43)))
	if v432 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	if v419 == int32(0) {
		v444 = v416
		goto L87
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	if v419 != 0 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v437 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v419))) = v437
	v447 = v437
	goto L86
L94:
	;
	v440 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v428)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v419))) = v440
	goto L96
L95:
	;
	goto L96
L96:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v428)+44))
	v444 = v442
	goto L87
L97:
	;
	if v448 <= int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v606 = v374 | base.B2i32(v395 == v416)
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+16)))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
	v609 = int32(0)
	v611 = v607 & base.B2i32(v608 != v609)
	if v396 < v609 {
		goto L113
	} else {
		goto L114
	}
L99:
	;
	v456 = v448 & int32(3)
	v457 = int32(4)
	v458 = v447 + v457
	if base.Ui32(v448) < base.Ui32(v457) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	if v456 == int32(0) {
		goto L98
	} else {
		goto L107
	}
L101:
	;
	v520 = v458
	v522 = int32(0)
	goto L100
L102:
	;
	goto L103
L103:
	;
	v465 = int32(0)
	v468 = v458
	v470 = v465
	v482 = v465
	goto L104
L104:
	;
	v493 = v451 + v470<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v493))) = v468
	v495 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v468)+6)))
	v496 = int32(8191)
	v498 = v468 + v495&v496
	*(*int32)(unsafe.Add(mBase, uint32(v493)+4)) = v498
	v500 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v498)+6)))
	v503 = v498 + v500&v496
	*(*int32)(unsafe.Add(mBase, uint32(v493)+8)) = v503
	v505 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v503)+6)))
	v508 = v503 + v505&v496
	*(*int32)(unsafe.Add(mBase, uint32(v493)+12)) = v508
	v510 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v508)+6)))
	v513 = v508 + v510&v496
	v514 = int32(4)
	v515 = v470 + v514
	v517 = v482 + v514
	if v517 != v448&int32(2147483644) {
		v468 = v513
		v470 = v515
		v482 = v517
		goto L104
	} else {
		goto L106
	}
L105:
	;
	v520 = v513
	v522 = v515
	goto L100
L106:
	;
	goto L105
L107:
	;
	v546 = v520
	v548 = v522
	v551 = int32(0)
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451+v548<<(uint(int32(2))%32)))) = v546
	v573 = int32(1)
	v575 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v546)+6)))
	v580 = v551 + v573
	if v580 != v456 {
		v546 = v546 + v575&int32(8191)
		v548 = v548 + v573
		v551 = v580
		goto L108
	} else {
		goto L110
	}
L109:
	;
	goto L98
L110:
	;
	goto L109
L111:
	;
	F_gistfillbuffer(m, v415, v451, v448, int32(1))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L10
	} else {
		goto L116
	}
L112:
	;
	F_PageInit(m, v629, int32(8192), int32(16))
	mBase = m.M
	v633 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v629)+16)))
	v634 = v629 + v633
	v635 = int32(65409)
	*(*uint16)(unsafe.Add(mBase, uint32(v634)+14)) = uint16(v635)
	*(*uint16)(unsafe.Add(mBase, uint32(v634)+12)) = uint16(v611)
	*(*int32)(unsafe.Add(mBase, uint32(v634)+8)) = int32(-1)
	goto L111
L113:
	;
	v615 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v615+(v396^int32(-1))<<(uint(int32(2))%32))))
	v629 = v621
	goto L112
L114:
	;
	goto L115
L115:
	;
	v623 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v629 = v623 + v396<<(uint(int32(13))%32) + int32(-8192)
	goto L112
L116:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
	if v643 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v415)+4)) = base.I32_wrap_i64(v356)
	*(*int32)(unsafe.Add(mBase, uint32(v415))) = base.I32_wrap_i64(int64(base.Ui64(v356) >> (uint(int64(32)) % 64)))
	F_MarkBufferDirty(m, v396)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L10
	} else {
		goto L129
	}
L118:
	;
	v646 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v415+v646)+8)) = int32(-1)
	v650 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415)+16)))
	v652 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v354)+12)))
	*(*uint32)(unsafe.Add(mBase, uint32(v415+v650))) = uint32(v652)
	v654 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415)+16)))
	v656 = *(*int64)(unsafe.Add(mBase, uint32(v354)+8))
	*(*uint32)(unsafe.Add(mBase, uint32(v415+v654)+4)) = uint32(v656)
	v658 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415)+16)))
	v659 = v415 + v658
	v660 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v659)+12)))
	v662 = v660 & int32(65527)
	*(*uint16)(unsafe.Add(mBase, uint32(v659)+12)) = uint16(v662)
	goto L117
L119:
	;
	goto L120
L120:
	;
	v664 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v354)+18)))
	if v370 < v664-int32(1) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v686 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415)+16)))
	v688 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v354)+12)))
	*(*uint32)(unsafe.Add(mBase, uint32(v415+v686))) = uint32(v688)
	v690 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415)+16)))
	v692 = *(*int64)(unsafe.Add(mBase, uint32(v354)+8))
	*(*uint32)(unsafe.Add(mBase, uint32(v415+v690)+4)) = uint32(v692)
	v694 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v354)+18)))
	v695 = int32(1)
	if (base.B2i32(v694-v695 <= v370)|v606)&v695 != 0 {
		goto L126
	} else {
		goto L127
	}
L122:
	;
	v672 = int32(0)
	F_XLogRecGetBlockTag(m, l0, (v370+int32(2))&int32(255), v672, v672, v27+int32(88))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L10
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v682 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415)+16)))
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v354)))
	*(*int32)(unsafe.Add(mBase, uint32(v415+v682)+8)) = v684
	goto L121
L125:
	;
	v678 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415)+16)))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v415+v678)+8)) = v680
	goto L121
L126:
	;
	v710 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415)+16)))
	v711 = v415 + v710
	v712 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v711)+12)))
	v714 = v712 & int32(65527)
	*(*uint16)(unsafe.Add(mBase, uint32(v711)+12)) = uint16(v714)
	goto L117
L127:
	;
	v701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+20)))
	if v701 != int32(1) {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v704 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415)+16)))
	v705 = v415 + v704
	v706 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v705)+12)))
	v708 = v706 | int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v705)+12)) = uint16(v708)
	goto L117
L129:
	;
	if v370 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v726 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v354)+18)))
	if base.Ui32(v386) < base.Ui32(v726) {
		v370 = v386
		v373 = v725
		v374 = v606
		goto L78
	} else {
		goto L135
	}
L131:
	;
	v725 = v396
	goto L130
L132:
	;
	goto L133
L133:
	;
	F_UnlockReleaseBuffer(m, v396)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L10
	} else {
		goto L134
	}
L134:
	;
	v725 = v373
	goto L130
L135:
	;
	goto L79
L136:
	;
	F_UnlockReleaseBuffer(m, v741)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L10
	} else {
		goto L140
	}
L137:
	;
	v756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v730)+76)))
	if v756 != int32(1) {
		goto L136
	} else {
		goto L138
	}
L138:
	;
	F_gistRedoClearFollowRight(m, l0, int32(0))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L10
	} else {
		goto L139
	}
L139:
	;
	goto L136
L140:
	;
	goto L4
L141:
	;
	if v769 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v773 = *(*int64)(unsafe.Add(mBase, uint32(v764)))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
	if v774 < int32(0) {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	goto L144
L144:
	;
	v815 = F_XLogReadBufferForRedo(m, l0, int32(1), v27+int32(76))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L10
	} else {
		goto L150
	}
L145:
	;
	v793 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v792)+16)))
	v794 = v793 + v792
	v795 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v794)+12)))
	v797 = v795 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v794)+12)) = uint16(v797)
	*(*int64)(unsafe.Add(mBase, uint32(v792)+24)) = v773
	v800 = int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v792)+12)) = uint16(v800)
	*(*uint32)(unsafe.Add(mBase, uint32(v792)+4)) = uint32(v765)
	v804 = int64(base.Ui64(v765) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v792))) = uint32(v804)
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
	F_MarkBufferDirty(m, v806)
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L10
	} else {
		goto L149
	}
L146:
	;
	v778 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v778+(v774^int32(-1))<<(uint(int32(2))%32))))
	v792 = v784
	goto L145
L147:
	;
	goto L148
L148:
	;
	v786 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v792 = v786 + v774<<(uint(int32(13))%32) + int32(-8192)
	goto L145
L149:
	;
	goto L144
L150:
	;
	if v815 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	if v819 < int32(0) {
		goto L155
	} else {
		goto L156
	}
L152:
	;
	goto L153
L153:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	if v848 != 0 {
		goto L160
	} else {
		goto L161
	}
L154:
	;
	v838 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v764)+8)))
	F_PageIndexTupleDelete(m, v837, v838)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L10
	} else {
		goto L158
	}
L155:
	;
	v823 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v823+(v819^int32(-1))<<(uint(int32(2))%32))))
	v837 = v829
	goto L154
L156:
	;
	goto L157
L157:
	;
	v831 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v837 = v831 + v819<<(uint(int32(13))%32) + int32(-8192)
	goto L154
L158:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v837))) = base.I64_rotr(v765, int64(32))
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	F_MarkBufferDirty(m, v844)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L10
	} else {
		goto L159
	}
L159:
	;
	goto L153
L160:
	;
	F_UnlockReleaseBuffer(m, v848)
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L10
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
	if v851 == int32(0) {
		goto L4
	} else {
		goto L164
	}
L163:
	;
	goto L162
L164:
	;
	F_UnlockReleaseBuffer(m, v851)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L10
	} else {
		goto L165
	}
L165:
	;
	goto L4
L166:
	;
	m.G0 = v27 + int32(96)
	return
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v109
	F_errmsg_internal(m, int32(169291), v27+int32(32))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L10
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(521581), int32(103), int32(441432))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L10
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v173
	F_errmsg_internal(m, int32(169291), v27+int32(16))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L10
	} else {
		goto L171
	}
L171:
	;
	F_errfinish(m, int32(521581), int32(139), int32(441432))
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L10
	} else {
		goto L172
	}
L172:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v38
	F_errmsg_internal(m, int32(57164), v27)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L10
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(521581), int32(430), int32(254532))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L10
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
