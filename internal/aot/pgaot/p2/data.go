package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateDataDirLockFile(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDataDirLockFile[0]))
	F_CreateLockFile(m, int32(_a_F_CreateDataDirLockFile_0), l0, int32(_a_F_CreateDataDirLockFile_1), int32(1), v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_dataExecPlaceToPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int64
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int64
	_ = v213
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	if l1 < int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_dataExecPlaceToPage[0]))
		v16 = int32(2)
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v13+(l1^int32(-1))<<(uint(v16)%32))))
		v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+16)))
		v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19+v20)+6)))
		if v22&v16 != 0 {
			v43 = v22
			v44 = v19
			v45 = v20
			v47 = v43 & int32(128)
			if v47 == int32(0) {
				v52 = v43 | int32(128)
				*(*uint16)(unsafe.Add(mBase, uint32(v44+v45)+6)) = uint16(v52)
				v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+16)))
				v56 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v44+v54)+4)) = uint16(v56)
			} else {
			}
			v59 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
			if base.B2i32(v59 == int32(0))|base.B2i32(l5 == v59) != 0 {
				v122 = int32(32)
			} else {
				v65 = int32(0)
				v72 = v44 + int32(32)
				v73 = base.B2i32(v47 == v65)
				v74 = v65
				v78 = v59
				for {
					v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+8)))
					v82 = base.B2i32(v79 != int32(0)) | v73
					if v79 != int32(1) {
						v85 = int32(1)
						v87 = int32(0)
						v89 = *(*int32)(unsafe.Add(mBase, uint32(v78)+20))
						v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+6)))
						v96 = (v90+v85)&int32(_a_F_dataExecPlaceToPage_0) + int32(8)
						if base.B2i32(v82&v85 == v87)|base.B2i32(v96 == v87) == v87 {
							base.MemoryCopy(m, v72, v89, v96)
						} else {
						}
						v105 = v72 + v96
						v106 = v74 + v96
					} else {
						v105 = v72
						v106 = v74
					}
					v109 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
					if v109 != l5 {
						v72 = v105
						v73 = v82
						v74 = v106
						v78 = v109
						continue
					} else {
						break
					}
					break
				}
				v122 = v106 + int32(32)
			}
			*(*uint16)(unsafe.Add(mBase, uint32(v44)+12)) = uint16(v122)
			F_MarkBufferDirty(m, l1)
			mBase = m.M
			v125 = m.ExcPending
			if v125 != 0 {
				return
			} else {
				v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+48))
				v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+118)))
				if v128 != int32(112) {
					return
				} else {
					v132 = *(*int32)(unsafe.Add(mBase, _c_F_dataExecPlaceToPage[1]))
					if v132 <= int32(0) {
						v135 = *(*int32)(unsafe.Add(mBase, uint32(v126)+32))
						if v135 != 0 {
							return
						} else {
							v136 = *(*int32)(unsafe.Add(mBase, uint32(v126)+40))
							if v136 != 0 {
								return
							} else {
								v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
								if v137 != 0 {
									return
								} else {
									F_XLogRegisterBuffer(m, int32(0), l1, int32(8))
									mBase = m.M
									v141 = m.ExcPending
									if v141 != 0 {
										return
									} else {
										v143 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
										v144 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
										F_XLogRegisterBufData(m, int32(0), v143, v144)
										mBase = m.M
										v146 = m.ExcPending
										if v146 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					} else {
						v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
						if v137 != 0 {
							return
						} else {
							F_XLogRegisterBuffer(m, int32(0), l1, int32(8))
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return
							} else {
								v143 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
								v144 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
								F_XLogRegisterBufData(m, int32(0), v143, v144)
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		} else {
			v148 = v19
			v149 = v20
			v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
			v153 = v148 + v150*int32(10)
			v155 = v153 + int32(22)
			*(*int32)(unsafe.Add(mBase, uint32(v155))) = base.I32_rotr(l4, int32(16))
			v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v148+v149)+4)))
			if v150 == int32(0) {
				v181 = v148 + v160*int32(10) + int32(32)
				v182 = v149
			} else {
				if v160+int32(1) == v150 {
					v181 = v155
					v182 = v149
				} else {
					v172 = int32(10)
					v175 = (v160-v150)*v172 + v172
					if v175 != 0 {
						base.MemoryCopy(m, v153+int32(32), v155, v175)
					} else {
					}
					v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v148)+16)))
					v181 = v155
					v182 = v179
				}
			}
			v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+8)))
			*(*uint16)(unsafe.Add(mBase, uint32(v181)+8)) = uint16(v183)
			v185 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
			*(*int64)(unsafe.Add(mBase, uint32(v181))) = v185
			v189 = v160 + int32(1)
			*(*uint16)(unsafe.Add(mBase, uint32(v148+v182)+4)) = uint16(v189)
			v194 = v160*int32(10) + int32(42)
			*(*uint16)(unsafe.Add(mBase, uint32(v148)+12)) = uint16(v194)
			F_MarkBufferDirty(m, l1)
			mBase = m.M
			v197 = m.ExcPending
			if v197 != 0 {
				return
			} else {
				v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+48))
				v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+118)))
				if v200 != int32(112) {
					return
				} else {
					v204 = *(*int32)(unsafe.Add(mBase, _c_F_dataExecPlaceToPage[1]))
					if v204 <= int32(0) {
						v207 = *(*int32)(unsafe.Add(mBase, uint32(v198)+32))
						if v207 != 0 {
							return
						} else {
							v208 = *(*int32)(unsafe.Add(mBase, uint32(v198)+40))
							if v208 != 0 {
								return
							} else {
								v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
								if v209 != 0 {
									return
								} else {
									*(*uint16)(unsafe.Add(mBase, _c_F_dataExecPlaceToPage[2])) = uint16(v150)
									v213 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
									*(*int64)(unsafe.Add(mBase, _c_F_dataExecPlaceToPage[3])) = v213
									v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+8)))
									*(*uint16)(unsafe.Add(mBase, _c_F_dataExecPlaceToPage[4])) = uint16(v216)
									F_XLogRegisterBuffer(m, int32(0), l1, int32(8))
									mBase = m.M
									v221 = m.ExcPending
									if v221 != 0 {
										return
									} else {
										F_XLogRegisterBufData(m, int32(0), int32(_a_F_dataExecPlaceToPage_1), int32(12))
										mBase = m.M
										v226 = m.ExcPending
										if v226 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					} else {
						v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
						if v209 != 0 {
							return
						} else {
							*(*uint16)(unsafe.Add(mBase, _c_F_dataExecPlaceToPage[2])) = uint16(v150)
							v213 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
							*(*int64)(unsafe.Add(mBase, _c_F_dataExecPlaceToPage[3])) = v213
							v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+8)))
							*(*uint16)(unsafe.Add(mBase, _c_F_dataExecPlaceToPage[4])) = uint16(v216)
							F_XLogRegisterBuffer(m, int32(0), l1, int32(8))
							mBase = m.M
							v221 = m.ExcPending
							if v221 != 0 {
								return
							} else {
								F_XLogRegisterBufData(m, int32(0), int32(_a_F_dataExecPlaceToPage_1), int32(12))
								mBase = m.M
								v226 = m.ExcPending
								if v226 != 0 {
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
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, _c_F_dataExecPlaceToPage[5]))
		v29 = v26 + l1<<(uint(int32(13))%32)
		v31 = v29 + int32(-8192)
		v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29-int32(_a_F_dataExecPlaceToPage_2)))))
		v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+v34-int32(_a_F_dataExecPlaceToPage_3)))))
		if v38&int32(2) == int32(0) {
			v148 = v31
			v149 = v34
			v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
			v153 = v148 + v150*int32(10)
			v155 = v153 + int32(22)
			*(*int32)(unsafe.Add(mBase, uint32(v155))) = base.I32_rotr(l4, int32(16))
			v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v148+v149)+4)))
			if v150 == int32(0) {
				v181 = v148 + v160*int32(10) + int32(32)
				v182 = v149
			} else {
				if v160+int32(1) == v150 {
					v181 = v155
					v182 = v149
				} else {
					v172 = int32(10)
					v175 = (v160-v150)*v172 + v172
					if v175 != 0 {
						base.MemoryCopy(m, v153+int32(32), v155, v175)
					} else {
					}
					v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v148)+16)))
					v181 = v155
					v182 = v179
				}
			}
			v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+8)))
			*(*uint16)(unsafe.Add(mBase, uint32(v181)+8)) = uint16(v183)
			v185 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
			*(*int64)(unsafe.Add(mBase, uint32(v181))) = v185
			v189 = v160 + int32(1)
			*(*uint16)(unsafe.Add(mBase, uint32(v148+v182)+4)) = uint16(v189)
			v194 = v160*int32(10) + int32(42)
			*(*uint16)(unsafe.Add(mBase, uint32(v148)+12)) = uint16(v194)
			F_MarkBufferDirty(m, l1)
			mBase = m.M
			v197 = m.ExcPending
			if v197 != 0 {
				return
			} else {
				v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+48))
				v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+118)))
				if v200 != int32(112) {
					return
				} else {
					v204 = *(*int32)(unsafe.Add(mBase, _c_F_dataExecPlaceToPage[1]))
					if v204 <= int32(0) {
						v207 = *(*int32)(unsafe.Add(mBase, uint32(v198)+32))
						if v207 != 0 {
							return
						} else {
							v208 = *(*int32)(unsafe.Add(mBase, uint32(v198)+40))
							if v208 != 0 {
								return
							} else {
								v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
								if v209 != 0 {
									return
								} else {
									*(*uint16)(unsafe.Add(mBase, _c_F_dataExecPlaceToPage[2])) = uint16(v150)
									v213 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
									*(*int64)(unsafe.Add(mBase, _c_F_dataExecPlaceToPage[3])) = v213
									v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+8)))
									*(*uint16)(unsafe.Add(mBase, _c_F_dataExecPlaceToPage[4])) = uint16(v216)
									F_XLogRegisterBuffer(m, int32(0), l1, int32(8))
									mBase = m.M
									v221 = m.ExcPending
									if v221 != 0 {
										return
									} else {
										F_XLogRegisterBufData(m, int32(0), int32(_a_F_dataExecPlaceToPage_1), int32(12))
										mBase = m.M
										v226 = m.ExcPending
										if v226 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					} else {
						v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
						if v209 != 0 {
							return
						} else {
							*(*uint16)(unsafe.Add(mBase, _c_F_dataExecPlaceToPage[2])) = uint16(v150)
							v213 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
							*(*int64)(unsafe.Add(mBase, _c_F_dataExecPlaceToPage[3])) = v213
							v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+8)))
							*(*uint16)(unsafe.Add(mBase, _c_F_dataExecPlaceToPage[4])) = uint16(v216)
							F_XLogRegisterBuffer(m, int32(0), l1, int32(8))
							mBase = m.M
							v221 = m.ExcPending
							if v221 != 0 {
								return
							} else {
								F_XLogRegisterBufData(m, int32(0), int32(_a_F_dataExecPlaceToPage_1), int32(12))
								mBase = m.M
								v226 = m.ExcPending
								if v226 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		} else {
			v43 = v38
			v44 = v31
			v45 = v34
			v47 = v43 & int32(128)
			if v47 == int32(0) {
				v52 = v43 | int32(128)
				*(*uint16)(unsafe.Add(mBase, uint32(v44+v45)+6)) = uint16(v52)
				v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+16)))
				v56 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v44+v54)+4)) = uint16(v56)
			} else {
			}
			v59 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
			if base.B2i32(v59 == int32(0))|base.B2i32(l5 == v59) != 0 {
				v122 = int32(32)
			} else {
				v65 = int32(0)
				v72 = v44 + int32(32)
				v73 = base.B2i32(v47 == v65)
				v74 = v65
				v78 = v59
				for {
					v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+8)))
					v82 = base.B2i32(v79 != int32(0)) | v73
					if v79 != int32(1) {
						v85 = int32(1)
						v87 = int32(0)
						v89 = *(*int32)(unsafe.Add(mBase, uint32(v78)+20))
						v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+6)))
						v96 = (v90+v85)&int32(_a_F_dataExecPlaceToPage_0) + int32(8)
						if base.B2i32(v82&v85 == v87)|base.B2i32(v96 == v87) == v87 {
							base.MemoryCopy(m, v72, v89, v96)
						} else {
						}
						v105 = v72 + v96
						v106 = v74 + v96
					} else {
						v105 = v72
						v106 = v74
					}
					v109 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
					if v109 != l5 {
						v72 = v105
						v73 = v82
						v74 = v106
						v78 = v109
						continue
					} else {
						break
					}
					break
				}
				v122 = v106 + int32(32)
			}
			*(*uint16)(unsafe.Add(mBase, uint32(v44)+12)) = uint16(v122)
			F_MarkBufferDirty(m, l1)
			mBase = m.M
			v125 = m.ExcPending
			if v125 != 0 {
				return
			} else {
				v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+48))
				v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+118)))
				if v128 != int32(112) {
					return
				} else {
					v132 = *(*int32)(unsafe.Add(mBase, _c_F_dataExecPlaceToPage[1]))
					if v132 <= int32(0) {
						v135 = *(*int32)(unsafe.Add(mBase, uint32(v126)+32))
						if v135 != 0 {
							return
						} else {
							v136 = *(*int32)(unsafe.Add(mBase, uint32(v126)+40))
							if v136 != 0 {
								return
							} else {
								v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
								if v137 != 0 {
									return
								} else {
									F_XLogRegisterBuffer(m, int32(0), l1, int32(8))
									mBase = m.M
									v141 = m.ExcPending
									if v141 != 0 {
										return
									} else {
										v143 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
										v144 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
										F_XLogRegisterBufData(m, int32(0), v143, v144)
										mBase = m.M
										v146 = m.ExcPending
										if v146 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					} else {
						v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
						if v137 != 0 {
							return
						} else {
							F_XLogRegisterBuffer(m, int32(0), l1, int32(8))
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return
							} else {
								v143 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
								v144 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
								F_XLogRegisterBufData(m, int32(0), v143, v144)
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
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
func F_process_data_packets(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int64
	_ = v263
	var v266 int32
	_ = v266
	var v267 int64
	_ = v267
	var v268 int64
	_ = v268
	var v269 int64
	_ = v269
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v390 int32
	_ = v390
	var v405 int32
	_ = v405
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v587 int32
	_ = v587
	v6 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(1120)
	m.G0 = v19
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v6
	v29 = v6
	v31 = v6
	goto L1
L1:
	;
	v44 = F_pgp_parse_pkt_hdr(m, l2, v19+int32(47), v19+int32(40), int32(1))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	m.G0 = v19 + int32(1120)
	return v587
L3:
	;
	goto L2
L4:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	F_pullf_free(m, v575)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L9
	} else {
		goto L179
	}
L5:
	;
	v564 = v546
	v565 = int32(0)
	v567 = int32(1)
	goto L4
L6:
	;
	v361 = F_pullf_read_max(m, v91, int32(4), v19+int32(1088), v19+int32(1084))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L9
	} else {
		goto L119
	}
L7:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	if v317 != 0 {
		goto L105
	} else {
		goto L106
	}
L8:
	;
	v313 = v310
	v316 = base.B2i32(v29 != int32(0))
	goto L7
L9:
	;
	return int32(0)
L10:
	;
	if v44 <= int32(0) {
		v310 = v44
		goto L8
	} else {
		goto L11
	}
L11:
	;
	if v29 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_px_debug(m, int32(_a_F_process_data_packets_0), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if v44 == int32(3) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v313 = int32(-100)
	v316 = int32(1)
	goto L7
L16:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+47)))
	switch v88 - int32(8) {
	case 0:
		goto L32
	default:
		goto L30
	case 3:
		goto L33
	case 11:
		goto L31
	}
L17:
	;
	v59 = l4
	goto L19
L18:
	;
	v59 = int32(0)
	goto L19
L19:
	;
	if v59 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v64 = F_palloc(m, int32(8))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L9
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v81 = F_pullf_create(m, v19+int32(36), int32(_a_F_process_data_packets_1), l0, l2)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L9
	} else {
		goto L27
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v44
	v71 = F_pullf_create(m, v19+int32(36), int32(_a_F_process_data_packets_2), v64, l2)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	if int32(0) <= v71 {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	F_pfree(m, v64)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v313 = v71
	v316 = int32(0)
	goto L7
L27:
	;
	if v81 < int32(0) {
		v310 = v81
		goto L8
	} else {
		goto L28
	}
L28:
	;
	goto L16
L29:
	;
	v564 = int32(-100)
	v565 = v307
	v567 = v308
	goto L4
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v88
	F_px_debug(m, int32(_a_F_process_data_packets_3), v19)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L9
	} else {
		goto L104
	}
L31:
	;
	if l4 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L32:
	;
	if l3 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L33:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	v94 = v19 + int32(48)
	v95 = F_pullf_read_fixed(m, v91, int32(1), v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	if v95 < int32(0) {
		v546 = v95
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+48)))
	v101 = F_pullf_read_fixed(m, v91, int32(1), v94)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	if v101 < int32(0) {
		v546 = v101
		goto L5
	} else {
		goto L37
	}
L37:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+48)))
	if v105 == int32(0) {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	v114 = v105
	goto L39
L39:
	;
	v126 = F_pullf_read(m, v91, v114, v19+int32(1088))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L9
	} else {
		goto L41
	}
L40:
	;
	F_px_debug(m, int32(_a_F_process_data_packets_4), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L9
	} else {
		goto L47
	}
L41:
	;
	if v126 < int32(0) {
		v546 = v126
		goto L5
	} else {
		goto L42
	}
L42:
	;
	if v126 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v130 = v114 - v126
	if v130 <= int32(0) {
		goto L6
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	goto L40
L46:
	;
	v114 = v130
	goto L39
L47:
	;
	v546 = int32(-100)
	goto L5
L48:
	;
	v140 = int32(0)
	F_px_debug(m, int32(_a_F_process_data_packets_5), v140)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L9
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if v31 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v307 = v140
	v308 = v31
	goto L29
L52:
	;
	v145 = int32(0)
	F_px_debug(m, int32(_a_F_process_data_packets_6), v145)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L9
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	v155 = F_pullf_read_fixed(m, v151, int32(1), v19+int32(48))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L9
	} else {
		goto L56
	}
L55:
	;
	v307 = v145
	v308 = int32(1)
	goto L29
L56:
	;
	if v155 < int32(0) {
		v546 = v155
		goto L5
	} else {
		goto L57
	}
L57:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v159
	switch v159 {
	case 0:
		goto L61
	case 1, 2:
		goto L60
	case 3:
		goto L59
	default:
		goto L58
	}
L58:
	;
	F_px_debug(m, int32(_a_F_process_data_packets_7), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L9
	} else {
		goto L72
	}
L59:
	;
	F_px_debug(m, int32(_a_F_process_data_packets_8), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L9
	} else {
		goto L67
	}
L60:
	;
	v167 = F_pgp_decompress_filter(m, v19+int32(48), l0, v151)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L9
	} else {
		goto L63
	}
L61:
	;
	v161 = int32(0)
	v163 = F_process_data_packets(m, l0, l1, v151, v161, v161)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	v546 = v163
	goto L5
L63:
	;
	if v167 < int32(0) {
		v546 = v167
		goto L5
	} else {
		goto L64
	}
L64:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	v172 = int32(0)
	v174 = F_process_data_packets(m, l0, l1, v171, v172, v172)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L9
	} else {
		goto L65
	}
L65:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	F_pullf_free(m, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L9
	} else {
		goto L66
	}
L66:
	;
	v546 = v174
	goto L5
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = int32(1)
	goto L68
L68:
	;
	v204 = F_pullf_read(m, v151, int32(_a_F_process_data_packets_9), v19+int32(48))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L9
	} else {
		goto L70
	}
L69:
	;
	v546 = v204
	goto L5
L70:
	;
	if int32(0) < v204 {
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v546 = int32(-100)
	goto L5
L73:
	;
	v215 = int32(0)
	F_px_debug(m, int32(_a_F_process_data_packets_10), v215)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L9
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v221 != 0 {
		v290 = int32(-12)
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v307 = v215
	v308 = v31
	goto L29
L77:
	;
	v298 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v298
	v564 = int32(0)
	v565 = v298
	v567 = v31
	goto L4
L78:
	;
	v564 = v290
	v565 = int32(0)
	v567 = v31
	goto L4
L79:
	;
	v222 = int32(-100)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	if v223 != int32(20) {
		v290 = v222
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(1)
	v234 = F_pullf_read_max(m, v226, int32(20), v19+int32(1084), v19+int32(1088))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L9
	} else {
		goto L81
	}
L81:
	;
	if v234 < int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v564 = v234
	v565 = int32(0)
	v567 = v31
	goto L4
L83:
	;
	goto L84
L84:
	;
	if v234 != int32(20) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	if v234 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v257 = v19 + int32(48)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v255)+16))
	m.T0[v258].(func(*base.Module, int32, int32))(m, v255, v257)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L9
	} else {
		goto L93
	}
L88:
	;
	F_px_debug(m, int32(_a_F_process_data_packets_11), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L9
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v234
	F_px_debug(m, int32(_a_F_process_data_packets_12), v19+int32(32))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L9
	} else {
		goto L92
	}
L91:
	;
	v564 = v222
	v565 = int32(0)
	v567 = v31
	goto L4
L92:
	;
	v564 = v222
	v565 = int32(0)
	v567 = v31
	goto L4
L93:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1084))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)+16))
	v263 = *(*int64)(unsafe.Add(mBase, uint32(v261)+8))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v19-int32(-64))))
	v267 = *(*int64)(unsafe.Add(mBase, uint32(v261)))
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v19)+56))
	v269 = *(*int64)(unsafe.Add(mBase, uint32(v19)+48))
	goto L95
L94:
	;
	goto L99
L95:
	;
	base.MemoryFill(m, v257, int32(0), int32(20))
	goto L97
L97:
	;
	goto L94
L98:
	;
	if base.I64_extend_i32_u(v262^v266)|(v267^v269|(v263^v268)) == int64(0) {
		goto L77
	} else {
		goto L102
	}
L99:
	;
	base.MemoryFill(m, v19+int32(1088), int32(0), int32(20))
	goto L101
L101:
	;
	goto L98
L102:
	;
	F_px_debug(m, int32(_a_F_process_data_packets_13), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L9
	} else {
		goto L103
	}
L103:
	;
	v290 = v222
	goto L78
L104:
	;
	v307 = int32(0)
	v308 = v31
	goto L29
L105:
	;
	F_pullf_free(m, v317)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L9
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v320 = int32(0)
	if v313 < v320 {
		v587 = v313
		goto L3
	} else {
		goto L109
	}
L108:
	;
	goto L107
L109:
	;
	if v31 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	F_px_debug(m, int32(_a_F_process_data_packets_14), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L9
	} else {
		goto L113
	}
L111:
	;
	v330 = v320
	goto L112
L112:
	;
	if base.B2i32(l4 == int32(0))|v316 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v330 = int32(-100)
	goto L112
L114:
	;
	v587 = v330
	goto L3
L115:
	;
	goto L116
L116:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v334 != 0 {
		v587 = v330
		goto L3
	} else {
		goto L117
	}
L117:
	;
	F_px_debug(m, int32(_a_F_process_data_packets_15), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L9
	} else {
		goto L118
	}
L118:
	;
	v587 = int32(-100)
	goto L3
L119:
	;
	if v361 != int32(4) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	F_px_debug(m, int32(_a_F_process_data_packets_4), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L9
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	goto L125
L123:
	;
	v546 = int32(-100)
	goto L5
L124:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v376 = int32(0)
	if base.B2i32(v375 == v376)|base.B2i32(base.Ui32(int32(-3)) < base.Ui32(v99-int32(118))) == v376 {
		goto L128
	} else {
		goto L129
	}
L125:
	;
	base.MemoryFill(m, v19+int32(1084), int32(0), int32(4))
	goto L127
L127:
	;
	goto L124
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v99
	F_px_debug(m, int32(_a_F_process_data_packets_16), v19+int32(16))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L9
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = base.B2i32(v99 == int32(117))
	v405 = int32(0)
	goto L132
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(1)
	goto L130
L132:
	;
	v416 = F_pullf_read(m, v91, int32(_a_F_process_data_packets_9), v19+int32(1088))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L9
	} else {
		goto L134
	}
L133:
	;
	if v416|base.B2i32(v405 == int32(0)) != 0 {
		v546 = v416
		goto L5
	} else {
		goto L177
	}
L134:
	;
	if int32(0) < v416 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v420 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L136:
	;
	goto L137
L137:
	;
	goto L133
L138:
	;
	if int32(0) <= v521 {
		v405 = v524
		goto L132
	} else {
		goto L176
	}
L139:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1088))
	v514 = F_mbuf_append(m, l1, v513, v416)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L9
	} else {
		goto L175
	}
L140:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v423 == int32(0) {
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1088))
	v427 = int32(0)
	if v405 == v427 {
		v437 = v427
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v438 = v416 + v426
	v445 = v426
	v447 = v437
	goto L146
L143:
	;
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426))))
	if v431 == int32(10) {
		v437 = int32(0)
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v434 = int32(13)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+48)) = uint8(v434)
	v437 = int32(1)
	goto L142
L145:
	;
	if int32(0) < v495 {
		goto L166
	} else {
		goto L167
	}
L146:
	;
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v445))))
	if v458 == int32(13) {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v495 = v488
	v496 = int32(0)
	goto L145
L148:
	;
	v461 = int32(1)
	v463 = v445 + v461
	if base.Ui32(v438) <= base.Ui32(v463) {
		v495 = v447
		v496 = v461
		goto L145
	} else {
		goto L151
	}
L149:
	;
	v473 = v445
	v474 = v458
	goto L150
L150:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(48)+v447))) = uint8(v474)
	v477 = v447 + int32(1)
	if v447 < int32(1023) {
		goto L159
	} else {
		goto L160
	}
L151:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v445)+1)))
	v467 = base.B2i32(v465 == int32(10))
	if v465 == int32(10) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v468 = v463
	goto L154
L153:
	;
	v468 = v445
	goto L154
L154:
	;
	if v465 == int32(10) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v471 = int32(10)
	goto L157
L156:
	;
	v471 = int32(13)
	goto L157
L157:
	;
	v473 = v468
	v474 = v471
	goto L150
L158:
	;
	v490 = v473 + int32(1)
	if base.Ui32(v490) < base.Ui32(v438) {
		v445 = v490
		v447 = v488
		goto L146
	} else {
		goto L164
	}
L159:
	;
	v488 = v477
	goto L158
L160:
	;
	goto L161
L161:
	;
	v480 = int32(0)
	v483 = F_mbuf_append(m, l1, v19+int32(48), v477)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L9
	} else {
		goto L162
	}
L162:
	;
	if v483 < int32(0) {
		v521 = v483
		v524 = v480
		goto L138
	} else {
		goto L163
	}
L163:
	;
	v488 = v480
	goto L158
L164:
	;
	goto L147
L165:
	;
	v521 = v512
	v524 = v496
	goto L138
L166:
	;
	v501 = F_mbuf_append(m, l1, v19+int32(48), v495)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L9
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v506 = int32(0)
	goto L172
L169:
	;
	if v501 < int32(0) {
		v512 = v501
		goto L165
	} else {
		goto L170
	}
L170:
	;
	goto L168
L171:
	;
	v512 = v506
	goto L165
L172:
	;
	base.MemoryFill(m, v19+int32(48), v506, int32(1024))
	goto L174
L174:
	;
	goto L171
L175:
	;
	v521 = v514
	v524 = v405
	goto L138
L176:
	;
	v546 = v521
	goto L5
L177:
	;
	v539 = F_mbuf_append(m, l1, int32(_a_F_process_data_packets_17), int32(1))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L9
	} else {
		goto L178
	}
L178:
	;
	v546 = v539
	goto L5
L179:
	;
	v578 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v578
	if v578 <= v564 {
		v29 = v565
		v31 = v567
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v587 = v564
	goto L3
}
