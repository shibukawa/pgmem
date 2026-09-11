package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RemoveXlogFile(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int64
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	v7 = m.G0
	v9 = v7 - int32(1072)
	m.G0 = v9
	v12 = l0 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v12
	v20 = F_pg_snprintf(m, v9+int32(48), int32(1024), int32(165223), v9+int32(32))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		v23 = int32(*(*uint8)(unsafe.Add(mBase, _consts[175])))
		if v23 != int32(1) {
			v75 = F_errstart(m, int32(13), int32(0))
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return
			} else {
				if v75 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v12
					F_errmsg_internal(m, int32(666413), v9)
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						F_errfinish(m, int32(467473), int32(4045), int32(366204))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							v89 = F_durable_unlink(m, v9+int32(48), int32(15))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								if v89 != 0 {
									m.G0 = v9 + int32(1072)
									return
								} else {
									v91 = int32(4321004)
									v93 = *(*int32)(unsafe.Add(mBase, _consts[176]))
									*(*int32)(unsafe.Add(mBase, _consts[176])) = v93 + int32(1)
									F_XLogArchiveCleanup(m, v12)
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return
									} else {
										m.G0 = v9 + int32(1072)
										return
									}
								}
							}
						}
					}
				} else {
					v89 = F_durable_unlink(m, v9+int32(48), int32(15))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						if v89 != 0 {
							m.G0 = v9 + int32(1072)
							return
						} else {
							v91 = int32(4321004)
							v93 = *(*int32)(unsafe.Add(mBase, _consts[176]))
							*(*int32)(unsafe.Add(mBase, _consts[176])) = v93 + int32(1)
							F_XLogArchiveCleanup(m, v12)
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return
							} else {
								m.G0 = v9 + int32(1072)
								return
							}
						}
					}
				}
			}
		} else {
			v26 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
			if base.Ui64(l1) < base.Ui64(v26) {
				v75 = F_errstart(m, int32(13), int32(0))
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return
				} else {
					if v75 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v12
						F_errmsg_internal(m, int32(666413), v9)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							F_errfinish(m, int32(467473), int32(4045), int32(366204))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								v89 = F_durable_unlink(m, v9+int32(48), int32(15))
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return
								} else {
									if v89 != 0 {
										m.G0 = v9 + int32(1072)
										return
									} else {
										v91 = int32(4321004)
										v93 = *(*int32)(unsafe.Add(mBase, _consts[176]))
										*(*int32)(unsafe.Add(mBase, _consts[176])) = v93 + int32(1)
										F_XLogArchiveCleanup(m, v12)
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return
										} else {
											m.G0 = v9 + int32(1072)
											return
										}
									}
								}
							}
						}
					} else {
						v89 = F_durable_unlink(m, v9+int32(48), int32(15))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return
						} else {
							if v89 != 0 {
								m.G0 = v9 + int32(1072)
								return
							} else {
								v91 = int32(4321004)
								v93 = *(*int32)(unsafe.Add(mBase, _consts[176]))
								*(*int32)(unsafe.Add(mBase, _consts[176])) = v93 + int32(1)
								F_XLogArchiveCleanup(m, v12)
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return
								} else {
									m.G0 = v9 + int32(1072)
									return
								}
							}
						}
					}
				}
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, _consts[3]))
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+320)))
				if v30 != int32(1) {
					v75 = F_errstart(m, int32(13), int32(0))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return
					} else {
						if v75 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v12
							F_errmsg_internal(m, int32(666413), v9)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return
							} else {
								F_errfinish(m, int32(467473), int32(4045), int32(366204))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									v89 = F_durable_unlink(m, v9+int32(48), int32(15))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										if v89 != 0 {
											m.G0 = v9 + int32(1072)
											return
										} else {
											v91 = int32(4321004)
											v93 = *(*int32)(unsafe.Add(mBase, _consts[176]))
											*(*int32)(unsafe.Add(mBase, _consts[176])) = v93 + int32(1)
											F_XLogArchiveCleanup(m, v12)
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return
											} else {
												m.G0 = v9 + int32(1072)
												return
											}
										}
									}
								}
							}
						} else {
							v89 = F_durable_unlink(m, v9+int32(48), int32(15))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								if v89 != 0 {
									m.G0 = v9 + int32(1072)
									return
								} else {
									v91 = int32(4321004)
									v93 = *(*int32)(unsafe.Add(mBase, _consts[176]))
									*(*int32)(unsafe.Add(mBase, _consts[176])) = v93 + int32(1)
									F_XLogArchiveCleanup(m, v12)
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return
									} else {
										m.G0 = v9 + int32(1072)
										return
									}
								}
							}
						}
					}
				} else {
					v37 = F_get_dirent_type(m, v9+int32(48), l0, int32(0), int32(13))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						if v37 != int32(2) {
							v75 = F_errstart(m, int32(13), int32(0))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return
							} else {
								if v75 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v12
									F_errmsg_internal(m, int32(666413), v9)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										F_errfinish(m, int32(467473), int32(4045), int32(366204))
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											v89 = F_durable_unlink(m, v9+int32(48), int32(15))
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return
											} else {
												if v89 != 0 {
													m.G0 = v9 + int32(1072)
													return
												} else {
													v91 = int32(4321004)
													v93 = *(*int32)(unsafe.Add(mBase, _consts[176]))
													*(*int32)(unsafe.Add(mBase, _consts[176])) = v93 + int32(1)
													F_XLogArchiveCleanup(m, v12)
													mBase = m.M
													v98 = m.ExcPending
													if v98 != 0 {
														return
													} else {
														m.G0 = v9 + int32(1072)
														return
													}
												}
											}
										}
									}
								} else {
									v89 = F_durable_unlink(m, v9+int32(48), int32(15))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										if v89 != 0 {
											m.G0 = v9 + int32(1072)
											return
										} else {
											v91 = int32(4321004)
											v93 = *(*int32)(unsafe.Add(mBase, _consts[176]))
											*(*int32)(unsafe.Add(mBase, _consts[176])) = v93 + int32(1)
											F_XLogArchiveCleanup(m, v12)
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return
											} else {
												m.G0 = v9 + int32(1072)
												return
											}
										}
									}
								}
							}
						} else {
							v44 = F_InstallXLogFileSegment(m, l2, v9+int32(48), int32(1), l1, l3)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								if v44 == int32(0) {
									v75 = F_errstart(m, int32(13), int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return
									} else {
										if v75 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = v12
											F_errmsg_internal(m, int32(666413), v9)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												F_errfinish(m, int32(467473), int32(4045), int32(366204))
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return
												} else {
													v89 = F_durable_unlink(m, v9+int32(48), int32(15))
													mBase = m.M
													v90 = m.ExcPending
													if v90 != 0 {
														return
													} else {
														if v89 != 0 {
															m.G0 = v9 + int32(1072)
															return
														} else {
															v91 = int32(4321004)
															v93 = *(*int32)(unsafe.Add(mBase, _consts[176]))
															*(*int32)(unsafe.Add(mBase, _consts[176])) = v93 + int32(1)
															F_XLogArchiveCleanup(m, v12)
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return
															} else {
																m.G0 = v9 + int32(1072)
																return
															}
														}
													}
												}
											}
										} else {
											v89 = F_durable_unlink(m, v9+int32(48), int32(15))
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return
											} else {
												if v89 != 0 {
													m.G0 = v9 + int32(1072)
													return
												} else {
													v91 = int32(4321004)
													v93 = *(*int32)(unsafe.Add(mBase, _consts[176]))
													*(*int32)(unsafe.Add(mBase, _consts[176])) = v93 + int32(1)
													F_XLogArchiveCleanup(m, v12)
													mBase = m.M
													v98 = m.ExcPending
													if v98 != 0 {
														return
													} else {
														m.G0 = v9 + int32(1072)
														return
													}
												}
											}
										}
									}
								} else {
									v50 = F_errstart(m, int32(13), int32(0))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										if v50 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v12
											F_errmsg_internal(m, int32(666483), v9+int32(16))
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return
											} else {
												F_errfinish(m, int32(467473), int32(4033), int32(366204))
												mBase = m.M
												v62 = m.ExcPending
												if v62 != 0 {
													return
												} else {
													v63 = int32(4321008)
													v65 = *(*int32)(unsafe.Add(mBase, _consts[177]))
													*(*int32)(unsafe.Add(mBase, _consts[177])) = v65 + int32(1)
													v69 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
													*(*int64)(unsafe.Add(mBase, uint32(l2))) = v69 + int64(1)
													F_XLogArchiveCleanup(m, v12)
													mBase = m.M
													v98 = m.ExcPending
													if v98 != 0 {
														return
													} else {
														m.G0 = v9 + int32(1072)
														return
													}
												}
											}
										} else {
											v63 = int32(4321008)
											v65 = *(*int32)(unsafe.Add(mBase, _consts[177]))
											*(*int32)(unsafe.Add(mBase, _consts[177])) = v65 + int32(1)
											v69 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
											*(*int64)(unsafe.Add(mBase, uint32(l2))) = v69 + int64(1)
											F_XLogArchiveCleanup(m, v12)
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return
											} else {
												m.G0 = v9 + int32(1072)
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
		}
	}
}
