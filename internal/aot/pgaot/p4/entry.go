package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cmpEntryAccumulator(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)))
	if v7 != v8 {
		if base.Ui32(v7) < base.Ui32(v8) {
			v13 = int32(-1)
		} else {
			v13 = int32(1)
		}
		return v13
	} else {
		v15 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+24)))
		v16 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+24)))
		if v15 != v16 {
			if v15 < v16 {
				v21 = int32(-1)
			} else {
				v21 = int32(1)
			}
			return v21
		} else {
			if v15 != 0 {
				v44 = int32(0)
				return v44
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v7<<(uint(int32(2))%32)+v24)+uint32(_c_F_cmpEntryAccumulator[0])))
				v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
				v37 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
				v38 = F_FunctionCall2Coll(m, v24+v7*int32(28)+int32(112), v35, v36, v37)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v44 = base.I32_wrap_i64(v38)
					return v44
				}
			}
		}
	}
}
func F_entryExecPlaceToPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	v5 = l4
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l1 < int32(0) {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_entryExecPlaceToPage[0]))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v16+(l1^int32(-1))<<(uint(int32(2))%32))))
		v30 = v22
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_entryExecPlaceToPage[1]))
		v30 = v24 + l1<<(uint(int32(13))%32) + int32(-8192)
	}
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
	if v32 == int32(1) {
		F_PageIndexTupleDelete(m, v30, v31)
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			if v5 == int32(-1) {
			} else {
				v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+16)))
				v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+v39)+6)))
				if v41&int32(2) != 0 {
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v30+v31<<(uint(int32(2))%32))+20))
					v50 = v30 + v47&int32(_a_F_entryExecPlaceToPage_0)
					v51 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)) = uint16(v51)
					*(*uint16)(unsafe.Add(mBase, uint32(v50)+2)) = uint16(v5)
					v55 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
					*(*uint16)(unsafe.Add(mBase, uint32(v50))) = uint16(v55)
				}
			}
			v58 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+6)))
			v63 = F_PageAddItemExtended(m, v30, v58, v59&int32(_a_F_entryExecPlaceToPage_1), v31, int32(0))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return
			} else {
				if v63 == v31 {
					F_MarkBufferDirty(m, l1)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+48))
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+118)))
						if v70 != int32(112) {
							m.G0 = v11 + int32(16)
							return
						} else {
							v74 = *(*int32)(unsafe.Add(mBase, _c_F_entryExecPlaceToPage[2]))
							if v74 <= int32(0) {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
								if v77 != 0 {
									m.G0 = v11 + int32(16)
									return
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v68)+40))
									if v78 != 0 {
										m.G0 = v11 + int32(16)
										return
									} else {
										v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
										if v79 != 0 {
											m.G0 = v11 + int32(16)
											return
										} else {
											v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
											*(*uint16)(unsafe.Add(mBase, _c_F_entryExecPlaceToPage[3])) = uint16(v31)
											*(*uint8)(unsafe.Add(mBase, _c_F_entryExecPlaceToPage[4])) = uint8(v80)
											F_XLogRegisterBuffer(m, int32(0), l1, int32(8))
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return
											} else {
												F_XLogRegisterBufData(m, int32(0), int32(_a_F_entryExecPlaceToPage_2), int32(4))
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return
												} else {
													v95 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
													v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95)+6)))
													F_XLogRegisterBufData(m, int32(0), v95, v96&int32(_a_F_entryExecPlaceToPage_1))
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return
													} else {
														m.G0 = v11 + int32(16)
														return
													}
												}
											}
										}
									}
								}
							} else {
								v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
								if v79 != 0 {
									m.G0 = v11 + int32(16)
									return
								} else {
									v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
									*(*uint16)(unsafe.Add(mBase, _c_F_entryExecPlaceToPage[3])) = uint16(v31)
									*(*uint8)(unsafe.Add(mBase, _c_F_entryExecPlaceToPage[4])) = uint8(v80)
									F_XLogRegisterBuffer(m, int32(0), l1, int32(8))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										F_XLogRegisterBufData(m, int32(0), int32(_a_F_entryExecPlaceToPage_2), int32(4))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											v95 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
											v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95)+6)))
											F_XLogRegisterBufData(m, int32(0), v95, v96&int32(_a_F_entryExecPlaceToPage_1))
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return
											} else {
												m.G0 = v11 + int32(16)
												return
											}
										}
									}
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return
					} else {
						v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v110 + int32(4)
						F_errmsg_internal(m, int32(_a_F_entryExecPlaceToPage_3), v11)
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_entryExecPlaceToPage_4), int32(571), int32(_a_F_entryExecPlaceToPage_5))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	} else {
		if v5 == int32(-1) {
		} else {
			v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+16)))
			v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+v39)+6)))
			if v41&int32(2) != 0 {
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v30+v31<<(uint(int32(2))%32))+20))
				v50 = v30 + v47&int32(_a_F_entryExecPlaceToPage_0)
				v51 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)) = uint16(v51)
				*(*uint16)(unsafe.Add(mBase, uint32(v50)+2)) = uint16(v5)
				v55 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
				*(*uint16)(unsafe.Add(mBase, uint32(v50))) = uint16(v55)
			}
		}
		v58 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
		v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+6)))
		v63 = F_PageAddItemExtended(m, v30, v58, v59&int32(_a_F_entryExecPlaceToPage_1), v31, int32(0))
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return
		} else {
			if v63 == v31 {
				F_MarkBufferDirty(m, l1)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+48))
					v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+118)))
					if v70 != int32(112) {
						m.G0 = v11 + int32(16)
						return
					} else {
						v74 = *(*int32)(unsafe.Add(mBase, _c_F_entryExecPlaceToPage[2]))
						if v74 <= int32(0) {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
							if v77 != 0 {
								m.G0 = v11 + int32(16)
								return
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v68)+40))
								if v78 != 0 {
									m.G0 = v11 + int32(16)
									return
								} else {
									v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
									if v79 != 0 {
										m.G0 = v11 + int32(16)
										return
									} else {
										v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
										*(*uint16)(unsafe.Add(mBase, _c_F_entryExecPlaceToPage[3])) = uint16(v31)
										*(*uint8)(unsafe.Add(mBase, _c_F_entryExecPlaceToPage[4])) = uint8(v80)
										F_XLogRegisterBuffer(m, int32(0), l1, int32(8))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return
										} else {
											F_XLogRegisterBufData(m, int32(0), int32(_a_F_entryExecPlaceToPage_2), int32(4))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												v95 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
												v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95)+6)))
												F_XLogRegisterBufData(m, int32(0), v95, v96&int32(_a_F_entryExecPlaceToPage_1))
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return
												} else {
													m.G0 = v11 + int32(16)
													return
												}
											}
										}
									}
								}
							}
						} else {
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
							if v79 != 0 {
								m.G0 = v11 + int32(16)
								return
							} else {
								v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
								*(*uint16)(unsafe.Add(mBase, _c_F_entryExecPlaceToPage[3])) = uint16(v31)
								*(*uint8)(unsafe.Add(mBase, _c_F_entryExecPlaceToPage[4])) = uint8(v80)
								F_XLogRegisterBuffer(m, int32(0), l1, int32(8))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									F_XLogRegisterBufData(m, int32(0), int32(_a_F_entryExecPlaceToPage_2), int32(4))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										v95 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
										v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95)+6)))
										F_XLogRegisterBufData(m, int32(0), v95, v96&int32(_a_F_entryExecPlaceToPage_1))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return
										} else {
											m.G0 = v11 + int32(16)
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v108 = m.ExcPending
				if v108 != 0 {
					return
				} else {
					v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v110 + int32(4)
					F_errmsg_internal(m, int32(_a_F_entryExecPlaceToPage_3), v11)
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_entryExecPlaceToPage_4), int32(571), int32(_a_F_entryExecPlaceToPage_5))
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_entryLocateLeafEntry(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int64
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	var v166 int32
	_ = v166
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v17 < int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v36 = int32(1)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
	if v37 == v36 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_entryLocateLeafEntry[0]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+(v17^int32(-1))<<(uint(int32(2))%32))))
	v35 = v27
	goto L1
L3:
	;
	goto L4
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_entryLocateLeafEntry[1]))
	v35 = v29 + v17<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	m.G0 = v15 + int32(16)
	return v166
L6:
	;
	v40 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v40)
	v166 = v36
	goto L5
L7:
	;
	goto L8
L8:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v42) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v56 = v48 + int32(1)
	if base.Ui32(int32(2)) <= base.Ui32(v56&int32(_a_F_entryLocateLeafEntry_0)) {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v48 = int32(base.Ui32(v42+int32(_a_F_entryLocateLeafEntry_1)) >> (uint(int32(2)) % 32))
	if v48&int32(_a_F_entryLocateLeafEntry_0) != 0 {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v52 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v52)
	v166 = int32(0)
	goto L5
L13:
	;
	goto L12
L14:
	;
	v66 = int32(1)
	v69 = v56
	goto L17
L15:
	;
	v155 = v56
	goto L16
L16:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v155)
	v166 = int32(0)
	goto L5
L17:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v82 = int32(base.Ui32((v69-v66)&int32(_a_F_entryLocateLeafEntry_2))>>(uint(int32(1))%32)) + v66
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v35+int32(20)+v82&int32(_a_F_entryLocateLeafEntry_0)<<(uint(int32(2))%32))))
	v91 = v35 + v88&int32(_a_F_entryLocateLeafEntry_3)
	v92 = F_gintuple_get_attrnum(m, v76, v91)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v155 = v141
	goto L16
L19:
	;
	return int32(0)
L20:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v99 = F_gintuple_get_key(m, v96, v91, v15+int32(15))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)))
	if v101 != v92 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v140 = base.B2i32(int32(0) < v136)
	if int32(0) < v136 {
		goto L40
	} else {
		goto L41
	}
L23:
	;
	if base.Ui32(v101) < base.Ui32(v92) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v107 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+64)))
	v108 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+15)))
	if v107 != v108 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v106 = int32(-1)
	goto L28
L27:
	;
	v106 = int32(1)
	goto L28
L28:
	;
	v136 = v106
	goto L22
L29:
	;
	if v107 < v108 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	if v107 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v113 = int32(-1)
	goto L34
L33:
	;
	v113 = int32(1)
	goto L34
L34:
	;
	v136 = v113
	goto L22
L35:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v92<<(uint(int32(2))%32)+v116)+uint32(_c_F_entryLocateLeafEntry[2])))
	v128 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v129 = F_FunctionCall2Coll(m, v116+v92*int32(28)+int32(112), v127, v128, v99)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L19
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v82)
	v166 = int32(1)
	goto L5
L38:
	;
	v131 = base.I32_wrap_i64(v129)
	if v131 != 0 {
		v136 = v131
		goto L22
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v141 = v69
	goto L42
L41:
	;
	v141 = v82
	goto L42
L42:
	;
	if int32(0) < v136 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v146 = v82 + int32(1)
	goto L45
L44:
	;
	v146 = v66
	goto L45
L45:
	;
	if base.Ui32(v146&int32(_a_F_entryLocateLeafEntry_0)) < base.Ui32(v141&int32(_a_F_entryLocateLeafEntry_0)) {
		v66 = v146
		v69 = v141
		goto L17
	} else {
		goto L46
	}
L46:
	;
	goto L18
}
